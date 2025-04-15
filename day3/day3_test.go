package day3_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day3"
)

func TestDelivery(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			">",
			2,
		},
		{
			"><><><><><><><><>",
			2,
		},
		{
			"^>v<",
			4,
		},
		{
			"^v^v^v^v^v",
			2,
		},
		{
			"^v><<>v^",
			5,
		},
		{
			getTestData(),
			2592,
		},
	}
	for _, test := range tests {
		if actual := day3.Deliver(test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRoboDelivery(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			"^v",
			3,
		},
		{
			"^>v<",
			3,
		},
		{
			"^v^v^v^v^v",
			11,
		},
		{
			"^v><<>v^",
			5,
		},
		{
			getTestData(),
			2360,
		},
	}
	for _, test := range tests {
		if actual := day3.RoboDeliver(test.input); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day3.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
