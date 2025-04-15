package day5_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day5"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{`jchzalrnumimnmhp
		haegwjzuvuyypxyu
		dvszwmarrgswjxmb`, 0},
		{`ugknbfddgicrmopn
		aaa
		jchzalrnumimnmhp
		haegwjzuvuyypxyu
		dvszwmarrgswjxmb`, 3},
		{getTestData(), 236},
	}
	for _, test := range tests {
		if actual := day5.Part1(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"qjhvhtzxzqqjkmpb\nxxyxx", 2},
		{"uurcxstgmygtbstg\nieodomkazucvgmuy", 0},
		{"qjhvhtzxzqqjkmpb\nxxyxx\nuurcxstgmygtbstg\nieodomkazucvgmuy", 2},
		{getTestData(), 51},
	}
	for _, test := range tests {
		if actual := day5.Part2(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day5.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
