package day12

import (
	"os"
	"testing"
)

// [1,2,3] and {"a":2,"b":4} both have a sum of 6.
// [[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
// {"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
// [] and {} both have a sum of 0.

func TestCount(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			"[1,2,3]",
			6,
		},
		{
			"{\"a\":2,\"b\":4}",
			6,
		},
		{
			"[11,22,33]",
			66,
		},
		{
			"[33,-11]",
			22,
		},
		{
			getTestData(),
			119433, // <-- Part 1
		},
	}
	for _, test := range test {
		actual := Count(test.input)
		if actual != test.expected {
			t.Errorf("expected %d, actual %d",
				test.expected, actual)
		}
	}
}

func TestCount2(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			sample,
			500,
		},
		{
			getTestData(),
			68466, // <-- Part 2
		},
	}
	for _, test := range test {
		actual := Count2(test.input)
		if actual != test.expected {
			t.Errorf("expected %d, actual %d",
				test.expected, actual)
		}
	}
}

const sample = `[
	[
	"red",
		[
		{
			"e": "green",
			"a": 100,
			"d": {
				"c": "violet",
				"a": "yellow",
				"b": "violet"
				},
			"c": "yellow",
			"h": "blue",
			"b": 100,
			"g": {
				"a": [
					"yellow",
					100,
					100,
					100,
					{"e": "violet","c": 123,"a": 101,"b": 87,"d": "red","f": 88}
				]
			}
		}
	]
]
]`

func getTestData() string {
	dat, err := os.ReadFile("../testdata/day12.txt")
	if err != nil {
		panic("not expected")
	}
	return string(dat)
}
