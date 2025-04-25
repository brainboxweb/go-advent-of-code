package day17_test

import (
	"testing"

	"github.com/brainboxweb/advent/day17"
	"github.com/stretchr/testify/require"
)

/*
The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to
move it into smaller containers. You take an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there
are four ways to do it:

15 and 10
20 and 5 (the first 5)
20 and 5 (the second 5)
15, 5, and 5
Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of eggnog?
*/

func TestGetCombos(t *testing.T) {
	var test = []struct {
		capacities []int
		volume     int
		expected   int
	}{
		{
			[]int{10, 5},
			15,
			1,
		},
		{
			[]int{20, 15, 10, 5, 5},
			25,
			4,
		},
		{
			day17data,
			150,
			4372,
		},
	}
	for _, test := range test {
		bb := day17.New(test.capacities)
		result := bb.GetCombos(test.volume)
		require.Equal(t, test.expected, result)
	}
}

func TestLowestCombo(t *testing.T) {
	var test = []struct {
		capacities []int
		volume     int
		expected   int
	}{
		{
			[]int{10, 5},
			15,
			2,
		},
		{
			[]int{20, 15, 10, 5, 5},
			25,
			2,
		},
		{
			day17data,
			150,
			4,
		},
	}
	for _, test := range test {
		bb := day17.New(test.capacities)
		result := bb.GetLowestCombo(test.volume)
		require.Equal(t, test.expected, result)
	}
}

var day17data = []int{11, 30, 47, 31, 32, 36, 3, 1, 5, 3, 32, 36, 15, 11, 46, 26, 28, 1, 19, 3}
