package day24_test

import (
	"testing"

	"github.com/brainboxweb/advent/day24"
)

// Damn. It's unstable
// 72050269


func TestRun(t *testing.T) {

	var tests = []struct {
		input      string
		groupCount int
		expected   int
	}{
		{
			`1
2
3
4
5
7
8
9
10
11`,
			3,
			99,
		},
		{
			day24data,
			3,
			10439961859, // <-- Part 1
		},
		{
			day24data,
			4,
			72050269, // <-- Part 2
		},
	}
	for _, test := range tests {
		if actual := day24.Run(test.input, test.groupCount); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day24data = `1
3
5
11
13
17
19
23
29
31
37
41
43
47
53
59
67
71
73
79
83
89
97
101
103
107
109
113`

//89,103,107,109 = 408
//23,41,43,47,53,59,73,79,101 = 519
//1,2,3,7,11,13,17,19,31,37,61,67,71,83,97
