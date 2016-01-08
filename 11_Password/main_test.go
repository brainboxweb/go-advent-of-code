package main

import (
	"testing"
)

//Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
//Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
//Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
//For example:

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

func TestAllowed(t *testing.T) {
	for _, test := range test {
		actual := allowed(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}

	}
}

var tests2 = []struct {
	input    string
	expected int
}{
	{"aaaa", 2},
	{"aaa", 1},
	{"abbceffg", 2},
	{"aaacccdss", 3},
}

func TestDoubleCount(t *testing.T) {
	for _, test := range tests2 {
		if actual := doubleCount(test.input, 0); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

//
//
//var tests2 = []struct {
//	input    string
//	expected bool
//}{
//	{"aaaa", true},
//	{"aaa", false},
//	{"abbceffg", true},
//}
//
//func TestDoublesNoOverlap(t *testing.T) {
//	for _, test := range tests2 {
//		if actual := doublesNoOverlap(test.input); actual != test.expected {
//			t.Errorf("Convert(%q) = %t, expected %t.",
//				test.input, actual, test.expected)
//		}
//	}
//}

var tests3 = []struct {
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

func TestRising(t *testing.T) {
	for _, test := range tests3 {
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

var tests4 = []struct {
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

func TestValid(t *testing.T) {
	for _, test := range tests4 {
		actual := valid(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}

	}
}

var tests5 = []struct {
	input    int
	expected string
}{
	{
		1,
		"b",
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

func TestBase26(t *testing.T) {
	for _, test := range tests5 {
		actual := base26(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s.",
				test.input, actual, test.expected)
		}

	}
}

var tests6 = []struct {
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

func TestBase26ToDecimal(t *testing.T) {
	for _, test := range tests6 {
		actual := base26ToDecimal(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}

	}
}

var tests7 = []struct {
	input    string
	expected string
}{
	{
		"hepxcrrq",
		"hepxxyzz",
	},
	{
		"hepxxyzz",
		"heqaabcc",
	},
}

func TestNextPassword(t *testing.T) {
	for _, test := range tests7 {
		actual := getNextPassword(test.input)
		if actual != test.expected {
			t.Errorf("Convert(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}

	}
}
