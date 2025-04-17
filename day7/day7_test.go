package day7

import (
	"os"
	"testing"

	"github.com/brainboxweb/advent/day7/wiring"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			getTestData("../testdata/day7.txt"),
			16076, // Part 1
		},
		{
			getTestData("../testdata/day7part2.txt"),
			2797, // Part 2
		},
	}
	for _, test := range tests {
		if actual := wiring.Run(test.input, "a"); actual != test.expected {
			t.Errorf("expected %d got %d",
				test.expected, actual)
		}
	}
}

func getTestData(fullpath string) string {
	dat, err := os.ReadFile(fullpath)
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
