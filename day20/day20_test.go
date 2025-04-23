package day20

import (
	"testing"
)

var tests = []struct {
	presents   int
	multiplier int
	limit      int
	expected   int
}{
	{
		30,
		10,
		-1,
		2,
	},
	{
		130,
		10,
		-1,
		8,
	},
	{
		120,
		10,
		-1,
		6,
	},
	{
		36000000,
		10,
		-1,
		831600,
	},
	{
		36000000,
		11,
		50,
		884520,
	},
}

func TestGetHouseNumbe(t *testing.T) {
	for _, test := range tests {
		if actual := GetHouseNumber(test.presents, test.multiplier, test.limit); actual != test.expected {
			t.Errorf("Convert(%d) = %d, expected %d.",
				test.presents, actual, test.expected)
		}
	}
}
