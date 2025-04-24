package day11

import (
	"testing"
)

//Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
//Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
//Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
//For example:

func TestAllowed(t *testing.T) {
	var test = []struct {
		input    string
		expected bool
	}{
		{
			"abc",
			true,
		},
		{
			"aib",
			false,
		},
		{
			"ao",
			false,
		},
		{
			"al",
			false,
		},
	}
	for _, test := range test {
		actual := allowed(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}

	}
}

func TestDoubleCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"aaaa", 2},
		{"aaa", 1},
		{"abbceffg", 2},
		{"aaacccdss", 3},
	}

	for _, test := range tests {
		if actual := doubleCount(test.input, 0); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRising(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{
			"abc",
			true,
		},
		{
			"cba",
			false,
		},
		{
			"sssxyzbb",
			true,
		},
	}
	for _, test := range tests {
		actual := rising(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}

	}
}

//hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
//abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
//abbcegjk fails the third requirement, because it only has one double letter (bb).
//The next password after abcdefgh is abcdffaa.
//The next password after ghijklmn is ghjaabcc, because you eventually skip all the passwords that start with ghi..., since i is not allowed.

func TestValid(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{
			"abbceffg",
			false,
		},
		{
			"abcdffaa",
			true,
		},
		{
			"ghjaabcc",
			true,
		},
	}
	for _, test := range tests {
		actual := valid(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}

	}
}

func TestBase26(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{

		{
			1,
			"b",
		},
		{
			2,
			"c",
		},
		{
			23,
			"x",
		},
		{
			24,
			"y",
		},
		{
			25,
			"z",
		},
		{
			26,
			"ba",
		},
		{
			286,
			"la",
		},
		{
			57647112526,
			"hepxcrrq",
		},
	}
	for _, test := range tests {
		actual := base26(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

func TestBase26ToDecimal(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			"b",
			1,
		},
		{
			"x",
			23,
		},
		{
			"y",
			24,
		},
		{
			"ba",
			26,
		},

		{
			"hepxcrrq",
			57647112526,
		},
	}
	for _, test := range tests {
		actual := base26ToDecimal(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}

	}
}

func TestNextPassword(t *testing.T) {
	var tests7 = []struct {
		input    string
		expected string
	}{
		{
			"cqjxjnds",
			"cqjxxyzz", // <-- Part 1
		},
		{
			"cqjxxyzz",
			"cqkaabcc", // <-- Part 2
		},
	}

	for _, test := range tests7 {
		actual := getNextPassword(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}

	}
}
