package naughty_test

import (
	"testing"

	"github.com/brainboxweb/advent/day5/naughty"
)

func TestNice(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"aeiss", true},
		{"aess", false},
		{"aeis", false},
		{"aeissab", false},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, test := range tests {
		phr := naughty.NewPhrase(test.input)
		if actual := phr.Nice(); actual != test.expected {
			t.Errorf("Convert(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

func TestNiceTwo(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, test := range tests {
		phr := naughty.NewPhrase(test.input)
		if actual := phr.NiceTwo(); actual != test.expected {
			t.Errorf("Convert(%s) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

