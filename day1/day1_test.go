package day1_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day1"
)

func TestFinalFloor(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"(", 1},
		{")", -1},
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
		{getTestData(), 280}, // Part 1 result
	}
	for _, test := range tests {
		if actual := day1.FinalFloor(test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestBasement(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()", 0},
		{"()()", 0},
		{"(())", 0},
		{")()()()()()()))(())", 1},
		{"())))))(((((((", 3},
		{getTestData(), 1797}, // Part 2 result
	}
	for _, test := range test {
		if actual, err := day1.Basement(test.input); actual != test.expected {
			if err != nil {
				continue
			}
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day1.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
