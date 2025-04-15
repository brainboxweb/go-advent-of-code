package main

import "testing"

var tests7 = []struct {
	input    string
	expected int
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
	{"bgvyzdsv", 254575},
}

func TestBitcoin5(t *testing.T) {
	for _, test := range tests7 {
		if actual := Bitcoin(5, test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

var tests8 = []struct {
	input    string
	expected int
}{
	{"bgvyzdsv", 1038736},
}

func TestBitcoin6(t *testing.T) {
	for _, test := range tests8 {
		if actual := Bitcoin(6, test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}
