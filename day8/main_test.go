package day8_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day8"
)

func TestPart1(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			`""
""`,
			4,
		},
		{
			`"abc"
"abc"`,
			4,
		},
		{
			getTestData(),
			1333, // <-- Part 1
		},
	}
	for _, test := range test {
		actual := day8.Part1(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			getTestData(),
			2046, // <-- Part 2
		},
	}
	for _, test := range test {
		actual := day8.Part2(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day8.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
