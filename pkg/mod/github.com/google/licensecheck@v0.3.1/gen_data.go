// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// This file generates data.gen.go.
// It embeds the text of all the licenses in the subdirectory "licenses"
// and constructs the data structures to represent them.
// Run by a "go:generate" comment in license.go.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/google/licensecheck"
)

var outFile = flag.String("o", "data.gen.go", "`file` to write")

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen: ")
	flag.Parse()

	filesLRE, err := filepath.Glob(filepath.Join("licenses", "*.lre"))
	if err != nil {
		log.Fatal(err)
	}
	if len(filesLRE) == 0 {
		log.Fatal("no license files")
	}

	code := outputTemplate
	out := new(bytes.Buffer)
	builtLRE := buildLRE(filesLRE)
	for _, file := range builtLRE {
		fmt.Fprintf(out, "\t\t{ID: %q, %s LRE: %v},\n", file.Name, file.Type, varName(file.Name+".lre"))
	}
	code = strings.Replace(code, "FILES_LIST", out.String(), -1)

	out.Reset()
	for _, file := range builtLRE {
		fmt.Fprintf(out, "const %s = `%s`\n",
			varName(file.Name+".lre"),
			bytes.ReplaceAll(file.Data, []byte("`"), []byte("` + \"`\" + `")))
	}
	code += out.String()

	src, err := format.Source([]byte(code))
	if err != nil {
		fd, err1 := ioutil.TempFile("", "license-data")
		if err1 == nil {
			_, err1 = fd.Write([]byte(code))
			if err1 == nil {
				log.Fatalf("parsing output (written to %s): %v", fd.Name(), err)
			}
			fd.Close()
		}
		log.Fatal("parsing output:", err)
	}
	err = ioutil.WriteFile(*outFile, src, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// varName returns the basename of the file, sanitized for use as a variable name,
// and given the prefix "license_".
func varName(file string) string {
	mapping := func(r rune) rune {
		if r == '-' || r == '.' || r == '+' {
			return '_'
		}
		return r
	}
	return "license_" + strings.Map(mapping, filepath.Base(file))
}

const outputTemplate = `
// Code generated by gen_data.go; DO NOT EDIT.

package licensecheck

var builtinLREs = []License{
	FILES_LIST
}
`

type fileData struct {
	Name string
	Type string
	Data []byte
}

func buildLRE(filesLRE []string) []fileData {
	var typ licensecheck.Type
	setType := func(s string) (string, error) {
		t, err := licensecheck.ParseType(s)
		if err != nil {
			return "", err
		}
		typ = t
		return "", nil
	}
	var out []fileData
	t := template.New("").Funcs(template.FuncMap{
		"list": templateList,
		"Type": setType,
	})
	t, err := t.ParseFiles(filesLRE...)
	if err != nil {
		log.Fatal("parsing LRE templates:", err)
	}
	for _, t := range t.Templates() {
		if strings.HasSuffix(t.Name(), ".lre") {
			var buf bytes.Buffer
			typ = licensecheck.Unknown
			if err := t.Execute(&buf, nil); err != nil {
				log.Fatalf("executing %s: %v", t.Name(), err)
			}
			if len(bytes.TrimSpace(buf.Bytes())) == 0 {
				// Only contained useful definitions.
				continue
			}
			tstr := ""
			if typ != licensecheck.Unknown {
				tstr = "Type: " + typ.String() + ","
			}
			out = append(out, fileData{strings.TrimSuffix(t.Name(), ".lre"), tstr, buf.Bytes()})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		ni, nj := out[i].Name, out[j].Name

		// Special case: BSD-4-Clause is a generalization of BSD-4-Clause-UC.
		// In case of multiple matches, licensecheck always returns the one earlier in the list.
		// Make BSD-4-Clause-UC the one earlier in the list.
		if strings.HasPrefix(ni, "BSD-4-Clause") && strings.HasPrefix(nj, "BSD-4-Clause") {
			ni, nj = nj, ni
		}

		// Special case: GFDL-1.[123]-invariants-* is a generalization of GFDL-1.[123]-no-invariants-*.
		// Reverse that order too.
		if strings.HasPrefix(ni, "GFDL-") && strings.HasPrefix(nj, "GFDL-") {
			ni, nj = nj, ni
		}

		return ni < nj
	})

	return out
}

// templateList returns xs, but it flattens any nested []interface{} into the main list.
// Called from templates as "list", to pass multiple arguments to templates.
func templateList(xs ...interface{}) []interface{} {
	var list []interface{}
	for _, x := range xs {
		switch x := x.(type) {
		case []interface{}:
			list = append(list, x...)
		default:
			list = append(list, x)
		}
	}
	return list
}
