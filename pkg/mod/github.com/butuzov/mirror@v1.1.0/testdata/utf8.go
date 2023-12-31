// Code generated by generate-tests; DO NOT EDIT.

package main

import (
	"unicode/utf8"
	. "unicode/utf8"
	pkg "unicode/utf8"
)


func main_utf8() {
	{
		
		_,_ = utf8.DecodeLastRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeLastRuneInString`
	}

	{
		
		_,_ = utf8.DecodeLastRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = DecodeLastRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeLastRuneInString`
	}

	{
		
		_,_ = DecodeLastRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = pkg.DecodeLastRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeLastRuneInString`
	}

	{
		
		_,_ = pkg.DecodeLastRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = utf8.DecodeLastRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeLastRune`
	}

	{
		
		_,_ = utf8.DecodeLastRuneInString("foobar") 
	}

	{
		
		_,_ = DecodeLastRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeLastRune`
	}

	{
		
		_,_ = DecodeLastRuneInString("foobar") 
	}

	{
		
		_,_ = pkg.DecodeLastRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeLastRune`
	}

	{
		
		_,_ = pkg.DecodeLastRuneInString("foobar") 
	}

	{
		
		_,_ = utf8.DecodeRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeRuneInString`
	}

	{
		
		_,_ = utf8.DecodeRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = DecodeRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeRuneInString`
	}

	{
		
		_,_ = DecodeRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = pkg.DecodeRune([]byte("foobar")) // want `avoid allocations with utf8\.DecodeRuneInString`
	}

	{
		
		_,_ = pkg.DecodeRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_,_ = utf8.DecodeRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeRune`
	}

	{
		
		_,_ = utf8.DecodeRuneInString("foobar") 
	}

	{
		
		_,_ = DecodeRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeRune`
	}

	{
		
		_,_ = DecodeRuneInString("foobar") 
	}

	{
		
		_,_ = pkg.DecodeRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.DecodeRune`
	}

	{
		
		_,_ = pkg.DecodeRuneInString("foobar") 
	}

	{
		
		_ = utf8.FullRune([]byte("foobar")) // want `avoid allocations with utf8\.FullRuneInString`
	}

	{
		
		_ = utf8.FullRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = FullRune([]byte("foobar")) // want `avoid allocations with utf8\.FullRuneInString`
	}

	{
		
		_ = FullRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = pkg.FullRune([]byte("foobar")) // want `avoid allocations with utf8\.FullRuneInString`
	}

	{
		
		_ = pkg.FullRune([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = utf8.FullRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.FullRune`
	}

	{
		
		_ = utf8.FullRuneInString("foobar") 
	}

	{
		
		_ = FullRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.FullRune`
	}

	{
		
		_ = FullRuneInString("foobar") 
	}

	{
		
		_ = pkg.FullRuneInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.FullRune`
	}

	{
		
		_ = pkg.FullRuneInString("foobar") 
	}

	{
		
		_ = utf8.RuneCount([]byte("foobar")) // want `avoid allocations with utf8\.RuneCountInString`
	}

	{
		
		_ = utf8.RuneCount([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = RuneCount([]byte("foobar")) // want `avoid allocations with utf8\.RuneCountInString`
	}

	{
		
		_ = RuneCount([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = pkg.RuneCount([]byte("foobar")) // want `avoid allocations with utf8\.RuneCountInString`
	}

	{
		
		_ = pkg.RuneCount([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = utf8.RuneCountInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.RuneCount`
	}

	{
		
		_ = utf8.RuneCountInString("foobar") 
	}

	{
		
		_ = RuneCountInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.RuneCount`
	}

	{
		
		_ = RuneCountInString("foobar") 
	}

	{
		
		_ = pkg.RuneCountInString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.RuneCount`
	}

	{
		
		_ = pkg.RuneCountInString("foobar") 
	}

	{
		
		_ = utf8.Valid([]byte("foobar")) // want `avoid allocations with utf8\.ValidString`
	}

	{
		
		_ = utf8.Valid([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = Valid([]byte("foobar")) // want `avoid allocations with utf8\.ValidString`
	}

	{
		
		_ = Valid([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = pkg.Valid([]byte("foobar")) // want `avoid allocations with utf8\.ValidString`
	}

	{
		
		_ = pkg.Valid([]byte{'f','o','o','b','a','r'}) 
	}

	{
		
		_ = utf8.ValidString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.Valid`
	}

	{
		
		_ = utf8.ValidString("foobar") 
	}

	{
		
		_ = ValidString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.Valid`
	}

	{
		
		_ = ValidString("foobar") 
	}

	{
		
		_ = pkg.ValidString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with utf8\.Valid`
	}

	{
		
		_ = pkg.ValidString("foobar") 
	}

}
