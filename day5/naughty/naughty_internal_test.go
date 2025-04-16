package naughty

import (
	"testing"
)

func TestDoubles(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"aa", true},
		{"ab", false},
		{"abb", true},
		{"ascdfggh", true},
		{"ascdfgigh", false},
	}
	for _, test := range tests {
		phr := NewPhrase(test.input)
		if actual := phr.doubleLetter(); actual != test.expected {
			t.Errorf("Convert(%q) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

func TestVowels(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"aixxxx", false},
		{"aixxxax", true},
		{"ozzzupppp", false},
		{"cooper", true},
	}
	for _, test := range tests {
		phr := NewPhrase(test.input)
		if actual := phr.vowels(); actual != test.expected {
			t.Errorf("Convert(%q) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

func TestBlacklist(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"ab", true},
		{"cd", true},
		{"pq", true},
		{"xy", true},
	}

	for _, test := range tests {
		phr := NewPhrase(test.input)
		if actual := phr.blacklist(); actual != test.expected {
			t.Errorf("Convert(%q) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

func TestDoublesNoOverlap(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"aaaa", true},
		{"aaa", false},
	}
	for _, test := range tests {
		phr := NewPhrase(test.input)
		if actual := phr.doublesNoOverlap(); actual != test.expected {
			t.Errorf("Convert(%q) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

func TestRepeatWithGap(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"axa", true},
		{"aa", false},
		{"abcdefeghi", true},
		{"xyx", true},
		{"aaa", true},
	}
	for _, test := range tests {
		phr := NewPhrase(test.input)
		if actual := phr.repeatWithGap(); actual != test.expected {
			t.Errorf("Convert(%q) = %t, expected %t.",
				test.input, actual, test.expected)
		}
	}
}

