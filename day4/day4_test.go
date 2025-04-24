package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent/day4"
)

func TestBitcoin5(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
		{"bgvyzdsv", 254575},
	}
	for _, test := range tests {
		if actual := day4.Bitcoin(5, test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestBitcoin6(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"bgvyzdsv", 1038736},
	}
	for _, test := range tests {
		if actual := day4.Bitcoin(6, test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}
