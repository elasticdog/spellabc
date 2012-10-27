// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

// Package spellalpha implements spelling alphabet code word encoding.
package spellalpha

import (
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

// codeWordFor returns the spelling alphabet code word for r.
// If r is not defined in the alphabet, returns r.
func codeWordFor(r rune) string {
	word, ok := alphabet[unicode.ToLower(r)]
	if !ok {
		return string(r)
	}

	switch {
	case unicode.IsLower(r):
		word = strings.ToLower(word)
	case unicode.IsUpper(r):
		word = strings.ToUpper(word)
	}
	return word
}

func encodeLine(s string) string {
	dst := make([]string, len(s))
	for index, char := range []rune(s) {
		dst[index] = codeWordFor(char)
	}
	return strings.Join(dst, " ")
}

// Encode converts a string into its spelling alphabet representation.
func Encode(s string) string {
	lines := strings.Split(s, "\n")
	dst := make([]string, len(lines))
	for index, line := range lines {
		dst[index] = encodeLine(line)
	}
	return strings.Join(dst, "\n")
}
