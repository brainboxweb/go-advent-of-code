package wiring_test

import (
	"testing"

	"github.com/brainboxweb/advent/day7/wiring"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			"123 -> a",
			123,
		},
		{
			"NOT xxx -> a\n123 -> xxx",
			65412,
		},
		{
			"123 -> x\n124 -> y\nx AND y -> a",
			120,
		},
		{
			"123 -> x\nx AND y -> a\n124 -> y",
			120,
		},
		{
			"111 -> x\n222 -> y\nx OR y -> a",
			255,
		},
		{
			"123 -> x\nx LSHIFT 1 -> a",
			246,
		},
		{
			"123 -> x\nx RSHIFT 1 -> a",
			61,
		},
		{
			"123 -> x\n456 -> y\nx AND y -> a",
			72,
		},
		{
			"222 -> y\nx AND y -> a\n111 -> x",
			78,
		},
		{
			one,
			65079,
		},
		{
			two,
			65079,
		},

		{
			three,
			65079,
		},
	}
	for _, test := range tests {
		if actual := wiring.Run(test.input, "a"); actual != test.expected {
			t.Errorf("expected %d got %d",
				test.expected, actual)
		}
	}
}

const one = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> a`

const two = `x AND y -> d
x OR y -> e
123 -> x
456 -> y
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> a`

const three = `xz AND yz -> dz
xz OR y -> ez
123 -> xz
456 -> yz
xz LSHIFT 2 -> fz
yz RSHIFT 2 -> gz
NOT xz -> hz
NOT yz -> az
x AND y -> d
x OR y -> e
123 -> x
456 -> y
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> a`
