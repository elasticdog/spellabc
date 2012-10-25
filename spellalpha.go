// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

// Package spellalpha implements spelling alphabet code word encoding.
package spellalpha

import (
	"fmt"
	"strings"
	"unicode"
)

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

// PrintPhonetic prints the phonetic representation of a string. Spaces
// are always added between code words and a newline is appended.
func PrintPhonetic(input string) {
	runes := []rune(input)
	for _, char := range runes[:len(runes)-1] {
		fmt.Printf("%s ", encode(char))
	}
	fmt.Printf("%s\n", encode(runes[len(runes)-1]))
}
