// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

/*
  phonetic converts characters into their equivalent code words using the
  NATO spelling alphabet.
*/

package main

import (
	"fmt"
	"github.com/gaal/go-options/options"
	"os"
	"strings"
	"unicode"
)

// mySpec is the options specification for processing command-line flags
const mySpec = `
phonetic -- convert characters into spelling alphabet code words

Usage: phonetic [OPTIONS] STRING...
--
v,verbose  include input characters along with output
h,help     display this help message
`

// alphabet is a map of characters and their associated code words
var alphabet = map[rune]string{
	'a':  "Alpha",
	'b':  "Bravo",
	'c':  "Charlie",
	'd':  "Delta",
	'e':  "Echo",
	'f':  "Foxtrot",
	'g':  "Golf",
	'h':  "Hotel",
	'i':  "India",
	'j':  "Juliet",
	'k':  "Kilo",
	'l':  "Lima",
	'm':  "Mike",
	'n':  "November",
	'o':  "Oscar",
	'p':  "Papa",
	'q':  "Quebec",
	'r':  "Romeo",
	's':  "Sierra",
	't':  "Tango",
	'u':  "Uniform",
	'v':  "Victor",
	'w':  "Whiskey",
	'x':  "X-ray",
	'y':  "Yankee",
	'z':  "Zulu",
	'0':  "Zero",
	'1':  "One",
	'2':  "Two",
	'3':  "Three",
	'4':  "Four",
	'5':  "Five",
	'6':  "Six",
	'7':  "Seven",
	'8':  "Eight",
	'9':  "Nine",
	' ':  "Space",
	'!':  "Exclamation",
	'"':  "DoubleQuote",
	'#':  "Hash",
	'$':  "Dollars",
	'%':  "Percent",
	'&':  "Ampersand",
	'(':  "LeftParens",
	')':  "RightParens",
	'*':  "Asterisk",
	'+':  "Plus",
	',':  "Comma",
	'-':  "Dash",
	'.':  "Period",
	'/':  "ForeSlash",
	':':  "Colon",
	';':  "SemiColon",
	'<':  "LessThan",
	'=':  "Equals",
	'>':  "GreaterThan",
	'?':  "Question",
	'@':  "At",
	'[':  "LeftBracket",
	'\'': "SingleQuote",
	'\\': "BackSlash",
	']':  "RightBracket",
	'^':  "Caret",
	'_':  "Underscore",
	'`':  "Backtick",
	'{':  "LeftBrace",
	'|':  "Pipe",
	'}':  "RightBrace",
	'~':  "Tilde",
}

// encode converts a character into its spelling alphabet code word
// representation (if it is known). For letters, it returns
// a lowercase/uppercase code word based on the capitalization.
func encode(char rune) string {
	word, ok := alphabet[unicode.ToLower(char)]
	if !ok {
		return string(char)
	}

	if unicode.IsLetter(char) {
		if unicode.IsLower(char) {
			word = strings.ToLower(word)
		} else {
			word = strings.ToUpper(word)
		}
	}
	return word
}

// printPhonetic prints the phonetic representation of a string. Spaces
// are always added between code words and a newline is appended.
func printPhonetic(input string) {
	runes := []rune(input)
	for _, char := range runes[:len(runes)-1] {
		fmt.Printf("%s ", encode(char))
	}
	fmt.Printf("%s\n", encode(runes[len(runes)-1]))
}

func main() {
	spec := options.NewOptions(mySpec)
	opt := spec.Parse(os.Args[1:])
	nonFlags := append(opt.Extra, opt.Leftover...)

	switch {
	case opt.GetBool("help"):
		spec.PrintUsageAndExit("")
	case len(nonFlags) < 1:
		spec.PrintUsageAndExit("You must supply a STRING to convert.")
	}

	for _, input := range nonFlags {
		if opt.GetBool("verbose") {
			fmt.Printf("%s  =  ", input)
		}
		printPhonetic(input)
	}
}
