package day25

import "testing"

func TestStep(t *testing.T) {
	var tests = []struct {
		x  int
		y  int
		xx int
		yy int
	}{
		{
			1,
			1,
			1,
			2,
		},
		{
			1,
			2,
			2,
			1,
		},
		{
			2,
			1,
			1,
			3,
		},
		{
			1,
			3,
			2,
			2,
		},
		{
			2,
			2,
			3,
			1,
		},
		{
			3,
			1,
			1,
			4,
		},
		{
			1,
			4,
			2,
			3,
		},
	}

	for _, test := range tests {
		if actualX, actualY := step(test.x, test.y); actualX != test.xx || actualY != test.yy {
			t.Errorf("Coord (%d, %d)) = (%d, %d), expected (%d, %d).",
				test.x, test.y, actualX, actualY, test.xx, test.yy)
		}
	}
}

func TestRun(t *testing.T) {

	var tests = []struct {
		start    int
		targetX  int
		targetY  int
		expected int
	}{
		{
			20151125,
			1,
			2,
			31916031,
		},
		{
			20151125,
			2,
			1,
			18749137,
		},
		{
			20151125,
			2,
			2,
			21629792,
		},
		{
			20151125,
			4,
			4,
			9380097,
		},
		{
			20151125,
			6,
			5,
			31663883,
		},
		{
			20151125,
			4,
			2,
			7726640,
		},
		//	To continue, please consult the code grid in the manual.  Enter the code at row 2978, column 3083.
		//	NB The rows and columns on my grid are transposed.
		{
			20151125,
			3083,
			2978,
			2650453, // Part 1 (there is no Part 2)
		},
	}

	for _, test := range tests {
		if actual := run(test.start, test.targetX, test.targetY); actual != test.expected {
			t.Errorf("Point (%d, %d) = %d, expected %d.",
				test.targetX, test.targetY, actual, test.expected)
		}
	}
}
