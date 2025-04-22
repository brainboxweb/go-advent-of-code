package day16_test

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day16"
)

func TestRun(t *testing.T) {

	var test = []struct {
		input    string
		known    string
		expected int
	}{
		{
			`Sue 1: goldfish: 9, cars: 0, samoyeds: 9
Sue 2: perfumes: 5, trees: 8, goldfish: 8`,
			`samoyeds: 9
pomeranians: 3`,
			1,
		},
		{
			fileToString("../testdata/day16.txt"),
			`children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`,
			40,
		},
	}
	for _, test := range test {
		if actual := day16.Part1(test.input, test.known, false); actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestRunAdvanced(t *testing.T) {
	var test = []struct {
		input    string
		known    string
		expected int
	}{
		{
			fileToString("../testdata/day16.txt"),
			`children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`,
			241,
		},
	}
	for _, test := range test {
		if actual := day16.Part1(test.input, test.known, true); actual != test.expected {
			t.Errorf("Parse(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func fileToString(fullpath string) string {
	b, err := os.ReadFile(fullpath)
	if err != nil {
		panic("not expected")
	}
	return string(b) // convert content to a 'string'

}
