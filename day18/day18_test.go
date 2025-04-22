package day18_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day18"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		loops    int
		expected int
	}{
		{
			`.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
			1,
			11,
		},
		{
			`.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
			2,
			8,
		},
		{
			`.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
			3,
			4,
		},
		{
			`.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
			4,
			4,
		},
		{
			getTestData("../testdata/day18.txt"),
			100,
			768, // <-- Part 1
		},
	}

	for _, test := range tests {
		if actual := day18.Run(test.input, test.loops); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRunOverrideCorners(t *testing.T) {
	var tests2 = []struct {
		input    string
		loops    int
		override bool
		expected int
	}{
		{
			`##.#.#
...##.
#....#
..#...
#.#..#
####.#`,
			1,
			true,
			18,
		},
		{
			`##.#.#
...##.
#....#
..#...
#.#..#
####.#`,
			2,
			true,
			18,
		},
		{
			getTestData("../testdata/day18.txt"),
			100,
			true,
			781, // <-- Part 2
		},
	}

	for _, test := range tests2 {
		if actual := day18.RunWithOverride(test.input, test.loops); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func getTestData(fullpath string) string {
	b, err := os.ReadFile(fullpath)
	if err != nil {
		panic("not expected")
	}

	return string(b)
}
