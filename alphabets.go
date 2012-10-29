// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

package spellabc

import ()

// encodeJan is an encoder alphabet for the Joint Army/Navy phonetic
// alphabet.
var encodeJan = map[rune]string{
	'a':  "Able",
	'b':  "Baker",
	'c':  "Charlie",
	'd':  "Dog",
	'e':  "Easy",
	'f':  "Fox",
	'g':  "George",
	'h':  "How",
	'i':  "Item",
	'j':  "Jig",
	'k':  "King",
	'l':  "Love",
	'm':  "Mike",
	'n':  "Nan",
	'o':  "Oboe",
	'p':  "Peter",
	'q':  "Queen",
	'r':  "Roger",
	's':  "Sugar",
	't':  "Tare",
	'u':  "Uncle",
	'v':  "Victor",
	'w':  "William",
	'x':  "X-ray",
	'y':  "Yoke",
	'z':  "Zebra",
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

// encodeLapd is an encoder alphabet for the LAPD phonetic alphabet.
// It is also known as the APCO phonetic alphabet.
var encodeLapd = map[rune]string{
	'a':  "Adam",
	'b':  "Boy",
	'c':  "Charles",
	'd':  "David",
	'e':  "Edward",
	'f':  "Frank",
	'g':  "George",
	'h':  "Henry",
	'i':  "Ida",
	'j':  "John",
	'k':  "King",
	'l':  "Lincoln",
	'm':  "Mary",
	'n':  "Nora",
	'o':  "Ocean",
	'p':  "Paul",
	'q':  "Queen",
	'r':  "Robert",
	's':  "Sam",
	't':  "Tom",
	'u':  "Union",
	'v':  "Victor",
	'w':  "William",
	'x':  "X-ray",
	'y':  "Young",
	'z':  "Zebra",
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

// encodeNato is an encoder alphabet for the NATO phonetic alphabet.
// It is also known as the International Radiotelephony Spelling Alphabet
// and the ICAO spelling alphabet.
var encodeNato = map[rune]string{
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

// encodeUsFinancial is an encoder alphabet for the US Financial phonetic
// alphabet.
var encodeUsFinancial = map[rune]string{
	'a':  "Adam",
	'b':  "Bob",
	'c':  "Carol",
	'd':  "David",
	'e':  "Eddie",
	'f':  "Frank",
	'g':  "George",
	'h':  "Harry",
	'i':  "Ike",
	'j':  "Jim",
	'k':  "Kenny",
	'l':  "Larry",
	'm':  "Mary",
	'n':  "Nancy",
	'o':  "Oliver",
	'p':  "Peter",
	'q':  "Quincy",
	'r':  "Roger",
	's':  "Sam",
	't':  "Thomas",
	'u':  "Uncle",
	'v':  "Vincent",
	'w':  "William",
	'x':  "Xavier",
	'y':  "Yogi",
	'z':  "Zachary",
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

// encodeWesternUnion is an encoder alphabet for the Western Union
// phonetic alphabet.
var encodeWesternUnion = map[rune]string{
	'a':  "Adams",
	'b':  "Boston",
	'c':  "Chicago",
	'd':  "Denver",
	'e':  "Easy",
	'f':  "Frank",
	'g':  "George",
	'h':  "Henry",
	'i':  "Ida",
	'j':  "John",
	'k':  "King",
	'l':  "Lincoln",
	'm':  "Mary",
	'n':  "NewYork",
	'o':  "Ocean",
	'p':  "Peter",
	'q':  "Queen",
	'r':  "Roger",
	's':  "Sugar",
	't':  "Thomas",
	'u':  "Union",
	'v':  "Victor",
	'w':  "William",
	'x':  "X-ray",
	'y':  "Young",
	'z':  "Zero",
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
