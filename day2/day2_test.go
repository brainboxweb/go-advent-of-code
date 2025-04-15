package day2_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day2"
)

func TestWrapping(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1x1x1", 7},
		{"1x1x1\n1x1x1", 14},
		{"2x1x1", 11},
		{"1x1x1\n2x1x1", 18},
		{"1x2x3", 24},
		{"3x2x1", 24},
		{"3x1x2", 24},
		{"2x3x4", 58},
		{"1x1x10", 43},
		{getTestData(), 1598415},
	}
	for _, test := range tests {
		if actual, _ := day2.Wrapping(test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRibbon(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
		{"1x1x1", 5},
		{"1x1x2", 6},
		{"2x1x1", 6},
		{"1x2x1", 6},
		{"2x3x4\n1x1x10", 48},
		{getTestData(), 3812909},
	}
	for _, test := range tests {
		if _, actual := day2.Wrapping(test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day2.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
