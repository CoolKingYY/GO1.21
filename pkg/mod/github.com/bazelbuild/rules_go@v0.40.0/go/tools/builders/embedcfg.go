// Copyright 2021 The Bazel Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// buildEmbedcfgFile writes an embedcfg file to be read by the compiler.
// An embedcfg file can be used in Go 1.16 or higher if the "embed" package
// is imported and there are one or more //go:embed comments in .go files.
// The embedcfg file maps //go:embed patterns to actual file names.
//
// The embedcfg file will be created in workDir, and its name is returned.
// The caller is responsible for deleting it. If no embedcfg file is needed,
// "" is returned with no error.
//
// All source files listed in goSrcs with //go:embed comments must be in one
// of the directories in embedRootDirs (not in a subdirectory). Embed patterns
// are evaluated relative to the source directory. Embed sources (embedSrcs)
// outside those directories are ignored, since they can't be matched by any
// valid pattern.
func buildEmbedcfgFile(goSrcs []fileInfo, embedSrcs, embedRootDirs []string, workDir string) (string, error) {
	// Check whether this package uses embedding and whether the toolchain
	// supports it (Go 1.16+). With Go 1.15 and lower, we'll try to compile
	// without an embedcfg file, and the compiler will complain the "embed"
	// package is missing.
	var major, minor int
	if n, err := fmt.Sscanf(runtime.Version(), "go%d.%d", &major, &minor); n != 2 || err != nil {
		// Can't parse go version. Maybe it's a development version; fall through.
	} else if major < 1 || (major == 1 && minor < 16) {
		return "", nil
	}
	importEmbed := false
	haveEmbed := false
	for _, src := range goSrcs {
		if len(src.embeds) > 0 {
			haveEmbed = true
			rootDir := findInRootDirs(src.filename, embedRootDirs)
			if rootDir == "" || strings.Contains(src.filename[len(rootDir)+1:], string(filepath.Separator)) {
				// Report an error if a source files appears in a subdirectory of
				// another source directory. In this situation, the same file could be
				// referenced with different paths.
				return "", fmt.Errorf("%s: source files with //go:embed should be in same directory. Allowed directories are:\n\t%s",
					src.filename,
					strings.Join(embedRootDirs, "\n\t"))
			}
		}
		for _, imp := range src.imports {
			if imp.path == "embed" {
				importEmbed = true
			}
		}
	}
	if !importEmbed || !haveEmbed {
		return "", nil
	}

	// Build a tree of embeddable files. This includes paths listed with
	// -embedsrc. If one of those paths is a directory, the tree includes
	// its files and subdirectories. Paths in the tree are relative to the
	// path in embedRootDirs that contains them.
	root, err := buildEmbedTree(embedSrcs, embedRootDirs)
	if err != nil {
		return "", err
	}

	// Resolve patterns to sets of files.
	var embedcfg struct {
		Patterns map[string][]string
		Files    map[string]string
	}
	embedcfg.Patterns = make(map[string][]string)
	embedcfg.Files = make(map[string]string)
	for _, src := range goSrcs {
		for _, embed := range src.embeds {
			matchedPaths, matchedFiles, err := resolveEmbed(embed, root)
			if err != nil {
				return "", err
			}
			embedcfg.Patterns[embed.pattern] = matchedPaths
			for i, rel := range matchedPaths {
				embedcfg.Files[rel] = matchedFiles[i]
			}
		}
	}

	// Write the configuration to a JSON file.
	embedcfgData, err := json.MarshalIndent(&embedcfg, "", "\t")
	if err != nil {
		return "", err
	}
	embedcfgName := filepath.Join(workDir, "embedcfg")
	if err := ioutil.WriteFile(embedcfgName, embedcfgData, 0o666); err != nil {
		return "", err
	}
	return embedcfgName, nil
}

// findInRootDirs returns a string from rootDirs which is a parent of the
// file path p. If there is no such string, findInRootDirs returns "".
func findInRootDirs(p string, rootDirs []string) string {
	dir := filepath.Dir(p)
	for _, rootDir := range rootDirs {
		if rootDir == dir ||
			(strings.HasPrefix(dir, rootDir) && len(dir) > len(rootDir)+1 && dir[len(rootDir)] == filepath.Separator) {
			return rootDir
		}
	}
	return ""
}

// embedNode represents an embeddable file or directory in a tree.
type embedNode struct {
	name       string                // base name
	path       string                // absolute file path
	children   map[string]*embedNode // non-nil for directory
	childNames []string              // sorted
}

// add inserts file nodes into the tree rooted at f for the slash-separated
// path src, relative to the absolute file path rootDir. If src points to a
// directory, add recursively inserts nodes for its contents. If a node already
// exists (for example, if a source file and a generated file have the same
// name), add leaves the existing node in place.
func (n *embedNode) add(rootDir, src string) error {
	// Create nodes for parents of src.
	parent := n
	parts := strings.Split(src, "/")
	for _, p := range parts[:len(parts)-1] {
		if parent.children[p] == nil {
			parent.children[p] = &embedNode{
				name:     p,
				children: make(map[string]*embedNode),
			}
		}
		parent = parent.children[p]
	}

	// Create a node for src. If src is a directory, recursively create nodes for
	// its contents. Go embedding ignores symbolic links, but Bazel may use links
	// for generated files and directories, so we follow them here.
	var visit func(*embedNode, string, os.FileInfo) error
	visit = func(parent *embedNode, path string, fi os.FileInfo) error {
		base := filepath.Base(path)
		if parent.children[base] == nil {
			parent.children[base] = &embedNode{name: base, path: path}
		}
		if !fi.IsDir() {
			return nil
		}
		node := parent.children[base]
		node.children = make(map[string]*embedNode)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		names, err := f.Readdirnames(0)
		f.Close()
		if err != nil {
			return err
		}
		for _, name := range names {
			cPath := filepath.Join(path, name)
			cfi, err := os.Stat(cPath)
			if err != nil {
				return err
			}
			if err := visit(node, cPath, cfi); err != nil {
				return err
			}
		}
		return nil
	}

	path := filepath.Join(rootDir, src)
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	return visit(parent, path, fi)
}

func (n *embedNode) isDir() bool {
	return n.children != nil
}

// get returns a tree node, given a slash-separated path relative to the
// receiver. get returns nil if no node exists with that path.
func (n *embedNode) get(path string) *embedNode {
	if path == "." || path == "" {
		return n
	}
	for _, part := range strings.Split(path, "/") {
		n = n.children[part]
		if n == nil {
			return nil
		}
	}
	return n
}

var errSkip = errors.New("skip")

// walk calls fn on each node in the tree rooted at n in depth-first pre-order.
func (n *embedNode) walk(fn func(rel string, n *embedNode) error) error {
	var visit func(string, *embedNode) error
	visit = func(rel string, node *embedNode) error {
		err := fn(rel, node)
		if err == errSkip {
			return nil
		} else if err != nil {
			return err
		}
		for _, name := range node.childNames {
			if err := visit(path.Join(rel, name), node.children[name]); err != nil && err != errSkip {
				return err
			}
		}
		return nil
	}
	err := visit("", n)
	if err == errSkip {
		return nil
	}
	return err
}

// buildEmbedTree constructs a logical directory tree of embeddable files.
// The tree may contain a mix of static and generated files from multiple
// root directories. Directory artifacts are recursively expanded.
func buildEmbedTree(embedSrcs, embedRootDirs []string) (root *embedNode, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("building tree of embeddable files in directories %s: %v", strings.Join(embedRootDirs, string(filepath.ListSeparator)), err)
		}
	}()

	// Add each path to the tree.
	root = &embedNode{name: "", children: make(map[string]*embedNode)}
	for _, src := range embedSrcs {
		rootDir := findInRootDirs(src, embedRootDirs)
		if rootDir == "" {
			// Embedded path cannot be matched by any valid pattern. Ignore.
			continue
		}
		rel := filepath.ToSlash(src[len(rootDir)+1:])
		if err := root.add(rootDir, rel); err != nil {
			return nil, err
		}
	}

	// Sort children in each directory node.
	var visit func(*embedNode)
	visit = func(node *embedNode) {
		node.childNames = make([]string, 0, len(node.children))
		for name, child := range node.children {
			node.childNames = append(node.childNames, name)
			visit(child)
		}
		sort.Strings(node.childNames)
	}
	visit(root)

	return root, nil
}

// resolveEmbed matches a //go:embed pattern in a source file to a set of
// embeddable files in the given tree.
func resolveEmbed(embed fileEmbed, root *embedNode) (matchedPaths, matchedFiles []string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%v: could not embed %s: %v", embed.pos, embed.pattern, err)
		}
	}()

	// Remove optional "all:" prefix from pattern and set matchAll flag if present.
	// See https://pkg.go.dev/embed#hdr-Directives for details.
	pattern := embed.pattern
	var matchAll bool
	if strings.HasPrefix(pattern, "all:") {
		matchAll = true
		pattern = pattern[4:]
	}

	// Check that the pattern has valid syntax.
	if _, err := path.Match(pattern, ""); err != nil || !validEmbedPattern(pattern) {
		return nil, nil, fmt.Errorf("invalid pattern syntax")
	}

	// Search for matching files.
	err = root.walk(func(matchRel string, matchNode *embedNode) error {
		if ok, _ := path.Match(pattern, matchRel); !ok {
			// Non-matching file or directory.
			return nil
		}

		// TODO: Should check that directories along path do not begin a new module
		// (do not contain a go.mod).
		// https://cs.opensource.google/go/go/+/master:src/cmd/go/internal/load/pkg.go;l=2158;drc=261fe25c83a94fc3defe064baed3944cd3d16959
		for dir := matchRel; len(dir) > 1; dir = filepath.Dir(dir) {
			if base := path.Base(matchRel); isBadEmbedName(base) {
				what := "file"
				if matchNode.isDir() {
					what = "directory"
				}
				if dir == matchRel {
					return fmt.Errorf("cannot embed %s %s: invalid name %s", what, matchRel, base)
				} else {
					return fmt.Errorf("cannot embed %s %s: in invalid directory %s", what, matchRel, base)
				}
			}
		}

		if !matchNode.isDir() {
			// Matching file. Add to list.
			matchedPaths = append(matchedPaths, matchRel)
			matchedFiles = append(matchedFiles, matchNode.path)
			return nil
		}

		// Matching directory. Recursively add all files in subdirectories.
		// Don't add hidden files or directories (starting with "." or "_"),
		// unless "all:" prefix was set.
		// See golang/go#42328.
		matchTreeErr := matchNode.walk(func(childRel string, childNode *embedNode) error {
			// TODO: Should check that directories along path do not begin a new module
			// https://cs.opensource.google/go/go/+/master:src/cmd/go/internal/load/pkg.go;l=2158;drc=261fe25c83a94fc3defe064baed3944cd3d16959
			if childRel != "" {
				base := path.Base(childRel)
				if isBadEmbedName(base) || (!matchAll && (strings.HasPrefix(base, ".") || strings.HasPrefix(base, "_"))) {
					if childNode.isDir() {
						return errSkip
					}
					return nil
				}
			}
			if !childNode.isDir() {
				matchedPaths = append(matchedPaths, path.Join(matchRel, childRel))
				matchedFiles = append(matchedFiles, childNode.path)
			}
			return nil
		})
		if matchTreeErr != nil {
			return matchTreeErr
		}
		return errSkip
	})
	if err != nil && err != errSkip {
		return nil, nil, err
	}
	if len(matchedPaths) == 0 {
		return nil, nil, fmt.Errorf("no matching files found")
	}
	return matchedPaths, matchedFiles, nil
}

func validEmbedPattern(pattern string) bool {
	return pattern != "." && fsValidPath(pattern)
}

// validPath reports whether the given path name
// is valid for use in a call to Open.
// Path names passed to open are unrooted, slash-separated
// sequences of path elements, like “x/y/z”.
// Path names must not contain a “.” or “..” or empty element,
// except for the special case that the root directory is named “.”.
//
// Paths are slash-separated on all systems, even Windows.
// Backslashes must not appear in path names.
//
// Copied from io/fs.ValidPath in Go 1.16beta1.
func fsValidPath(name string) bool {
	if name == "." {
		// special case
		return true
	}

	// Iterate over elements in name, checking each.
	for {
		i := 0
		for i < len(name) && name[i] != '/' {
			if name[i] == '\\' {
				return false
			}
			i++
		}
		elem := name[:i]
		if elem == "" || elem == "." || elem == ".." {
			return false
		}
		if i == len(name) {
			return true // reached clean ending
		}
		name = name[i+1:]
	}
}

// isBadEmbedName reports whether name is the base name of a file that
// can't or won't be included in modules and therefore shouldn't be treated
// as existing for embedding.
//
// TODO: This should use the equivalent of golang.org/x/mod/module.CheckFilePath instead of fsValidPath.
// https://cs.opensource.google/go/go/+/master:src/cmd/go/internal/load/pkg.go;l=2200;drc=261fe25c83a94fc3defe064baed3944cd3d16959
func isBadEmbedName(name string) bool {
	if !fsValidPath(name) {
		return true
	}
	switch name {
	// Empty string should be impossible but make it bad.
	case "":
		return true
	// Version control directories won't be present in module.
	case ".bzr", ".hg", ".git", ".svn":
		return true
	}
	return false
}
