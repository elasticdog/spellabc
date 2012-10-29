// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

package spellalpha

import (
	"testing"
)

type testpair struct {
	decoded, encoded string
}

var pairs = []testpair{
	{"", ""},
	{"h", "hotel"},
	{"H", "HOTEL"},
	{"He", "HOTEL echo"},
	{"Hel", "HOTEL echo lima"},
	{"Hell", "HOTEL echo lima lima"},
	{"Hello", "HOTEL echo lima lima oscar"},
	{"Hello World!", "HOTEL echo lima lima oscar Space WHISKEY oscar romeo lima delta Exclamation"},
	{"ABC\ndef", "ALPHA BRAVO CHARLIE\ndelta echo foxtrot"},
}

func testEqual(t *testing.T, msg string, args ...interface{}) bool {
	if args[len(args)-2] != args[len(args)-1] {
		t.Errorf(msg, args...)
		return false
	}
	return true
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		got := NatoAlphabet.Encode(p.decoded)
		testEqual(t, "Encode(%q) = %q, want %q", p.decoded, got, p.encoded)
	}
}
