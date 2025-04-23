package day23_test

import (
	"testing"

	"github.com/brainboxweb/advent/day23"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input           string
		intialValueForA int
		expectedB       int
	}{
		{
			day23data,
			0,
			170, // <-- Part 1
		},
		{
			day23data,
			1,
			247, // <-- Part 2
		},
	}

	for _, test := range tests {

		result := day23.Run(test.input, test.intialValueForA)

		if result != test.expectedB {
			t.Errorf("expected %d, actual %df",
				test.expectedB, result)
		}
	}
}

const day23data = `jio a, +16
inc a
inc a
tpl a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
tpl a
inc a
jmp +23
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`
