package day8

import (
	"testing"
)

/*
"" is 2 characters of code (the two double quotes), but the string contains zero characters.
"abc" is 5 characters of code, but 3 characters in the string data.
"aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
"\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.
*/

func TestStrings(t *testing.T) {
	var test = []struct {
		input     string
		expected  int
		expected2 int
	}{
		{
			`""`,
			2,
			0,
		},
		{
			`"abc"`,
			5,
			3,
		},
		{
			`"aaa\"aaa"`,
			10,
			7,
		},
		{
			`"\x27"`,
			6,
			1,
		},
	}
	for _, test := range test {
		actual, actual2 := stringInfo(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
		if actual2 != test.expected2 {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual2, test.expected2)
		}
	}
}

// "" encodes to "\"\"", an increase from 2 characters to 6.
// "abc" encodes to "\"abc\"", an increase from 5 characters to 9.
// "aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
// "\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.

func TestStrings2(t *testing.T) {
	var test = []struct {
		input     string
		expected  int
		expected2 int
	}{
		{
			`""`,
			2,
			6,
		},
		{
			`"abc"`,
			5,
			9,
		},
		{
			`"aaa\"aaa"`,
			10,
			16,
		},
		{
			`"\x27"`,
			6,
			11,
		},
	}
	for _, test := range test {
		actual, actual2 := stringInfo2(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
		if actual2 != test.expected2 {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual2, test.expected2)
		}
	}
}
