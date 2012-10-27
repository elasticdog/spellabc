// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

// Package spellalpha implements spelling alphabet code word encoding.
package spellalpha

import (
	"strings"
	"unicode"
)

// An Encoding is a character to code word mapping scheme, defined by a
// spelling alphabet. The most common is the "NATO" spelling alphabet.
type Encoding struct {
	alphabet map[rune]string
}

// NewEncoding returns a new Encoding defined by the given alphabet, which
// must a map of rune -> string pairs.
func NewEncoding(encoder map[rune]string) *Encoding {
	e := new(Encoding)
	e.alphabet = encoder
	return e
}

// JanAlphabet is the Joint Army/Navy phonetic alphabet.
var JanAlphabet = NewEncoding(encodeJan)

// NatoAlphabet is the NATO phonetic alphabet.
// It is also known as the International Radiotelephony Spelling Alphabet
// and the ICAO spelling alphabet.
var NatoAlphabet = NewEncoding(encodeNato)

// UsFinancialAlphabet is the US Financial phonetic alphabet.
var UsFinancialAlphabet = NewEncoding(encodeUsFinancial)

// WesternUnionAlphabet is the Western Union phonetic alphabet.
var WesternUnionAlphabet = NewEncoding(encodeWesternUnion)

// codeWordFor returns the spelling alphabet code word for r.
// If r is not defined in the alphabet, returns r.
func (enc *Encoding) codeWordFor(r rune) string {
	word, ok := enc.alphabet[unicode.ToLower(r)]
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

func (enc *Encoding) encodeLine(s string) string {
	dst := make([]string, len(s))
	for index, char := range []rune(s) {
		dst[index] = enc.codeWordFor(char)
	}
	return strings.Join(dst, " ")
}

// Encode converts s into its spelling alphabet representation, defined by enc.
func (enc *Encoding) Encode(s string) string {
	lines := strings.Split(s, "\n")
	dst := make([]string, len(lines))
	for index, line := range lines {
		dst[index] = enc.encodeLine(line)
	}
	return strings.Join(dst, "\n")
}
