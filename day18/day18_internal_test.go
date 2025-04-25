package day18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

	expected := [][]int{}

	row0 := []int{0, 1, 0, 1, 0, 1}
	row1 := []int{0, 0, 0, 1, 1, 0}
	row2 := []int{1, 0, 0, 0, 0, 1}
	row3 := []int{0, 0, 1, 0, 0, 0}
	row4 := []int{1, 0, 1, 0, 0, 1}
	row5 := []int{1, 1, 1, 1, 0, 0}

	expected = append(expected, row0)
	expected = append(expected, row1)
	expected = append(expected, row2)
	expected = append(expected, row3)
	expected = append(expected, row4)
	expected = append(expected, row5)

	result := parse(input)

	require.Equal(t, expected, result)
}
