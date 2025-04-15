package main

import (
	"reflect"
	"testing"
)

//1 becomes 11 (1 copy of digit 1).
//11 becomes 21 (2 copies of digit 1).
//21 becomes 1211 (one 2 followed by one 1).
//1211 becomes 111221 (one 1, one 2, and two 1s).
//111221 becomes 312211 (three 1s, two 2s, and one 1).

var test = []struct {
	input    []int
	expected []int
}{
	{
		[]int{1},
		[]int{1, 1},
	},
	{
		[]int{1, 1},
		[]int{2, 1},
	},
	{
		[]int{2, 1},
		[]int{1, 2, 1, 1},
	},
	{
		[]int{1, 2, 1, 1},
		[]int{1, 1, 1, 2, 2, 1},
	},
}

func TestSay(t *testing.T) {
	for _, test := range test {
		if actual := say(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Process(%q) = %q, expected %q.",
				test.input, actual, test.expected, len(actual))
		}
	}
}

var test2 = []struct {
	input    int
	loops    int
	expected int
}{
	{
		1,
		1,
		2,
	},
	{
		11,
		1,
		2,
	},
	{
		21,
		1,
		4,
	},
	{
		1211,
		1,
		6,
	},
	{
		3113322113,
		40,
		329356,
	},
	{
		3113322113,
		50,
		4666278,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		if actual := Run(test.input, test.loops); actual != test.expected {
			t.Errorf("Parse(%d) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}
