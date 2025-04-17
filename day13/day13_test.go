package day13_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day13"
)

func TestRun(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			sample,
			330,
		},
		{
			getTestData(),
			709,
		},
	}
	for _, test := range test {
		if actual := day13.Run(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRunWithGuest(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			sample,
			286,
		},
		{
			getTestData(),
			668,
		},
	}

	for _, test := range test {
		if actual := day13.RunWithGuest(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day13.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}

const sample = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`
