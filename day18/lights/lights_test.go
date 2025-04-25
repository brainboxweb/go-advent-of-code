package lights_test

import (
	"testing"

	"github.com/brainboxweb/advent/day18/lights"
	"github.com/stretchr/testify/require"
)

func TestSwitch(t *testing.T) {
	input := [][]int{}

	row0 := []int{1, 1, 0}
	row1 := []int{0, 1, 0}
	row2 := []int{0, 0, 0}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)

	ll := lights.New(input)

	expected := [][]int{}

	// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
	// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

	_row0 := []int{1, 1, 0}
	_row1 := []int{1, 1, 0}
	_row2 := []int{0, 0, 0}

	expected = append(expected, _row0)
	expected = append(expected, _row1)
	expected = append(expected, _row2)

	result := ll.SwitchLights()

	require.Equal(t, expected, result)
}

func TestSwitch2(t *testing.T) {
	input := [][]int{}

	row0 := []int{1, 1, 1}
	row1 := []int{1, 1, 1}
	row2 := []int{1, 1, 1}
	row3 := []int{1, 1, 1}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)
	input = append(input, row3)

	ll := lights.New(input)

	expected := [][]int{}

	// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
	// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

	_row0 := []int{1, 0, 1}
	_row1 := []int{0, 0, 0}
	_row2 := []int{0, 0, 0}
	_row3 := []int{1, 0, 1}

	expected = append(expected, _row0)
	expected = append(expected, _row1)
	expected = append(expected, _row2)
	expected = append(expected, _row3)

	result := ll.SwitchLights()

	require.Equal(t, expected, result)
}

func TestSwitch3(t *testing.T) {
	input := [][]int{}

	// 	.#.#.#
	// ...##.
	// #....#
	// ..#...
	// #.#..#
	// ####..

	row0 := []int{0, 1, 0, 1, 0, 1}
	row1 := []int{0, 0, 0, 1, 1, 0}
	row2 := []int{1, 0, 0, 0, 0, 1}
	row3 := []int{0, 0, 1, 0, 0, 0}
	row4 := []int{1, 0, 1, 0, 0, 1}
	row5 := []int{1, 1, 1, 1, 0, 0}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)
	input = append(input, row3)
	input = append(input, row4)
	input = append(input, row5)

	ll := lights.New(input)

	expected := [][]int{}
	// ..##..
	// ..##.#
	// ...##.
	// ......
	// #.....
	// #.##..
	// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
	// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

	_row0 := []int{0, 0, 1, 1, 0, 0}
	_row1 := []int{0, 0, 1, 1, 0, 1}
	_row2 := []int{0, 0, 0, 1, 1, 0}
	_row3 := []int{0, 0, 0, 0, 0, 0}
	_row4 := []int{1, 0, 0, 0, 0, 0}
	_row5 := []int{1, 0, 1, 1, 0, 0}

	expected = append(expected, _row0)
	expected = append(expected, _row1)
	expected = append(expected, _row2)
	expected = append(expected, _row3)
	expected = append(expected, _row4)
	expected = append(expected, _row5)

	result := ll.SwitchLights()

	require.Equal(t, expected, result)
}

func TestNeightbourOnCount(t *testing.T) {
	input := [][]int{}

	row0 := []int{0, 1, 0}
	row1 := []int{0, 0, 1}
	row2 := []int{1, 1, 1}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)

	ll := lights.New(input)

	result := ll.NeighbourOnCount(1, 1)
	expected := 5
	require.Equal(t, expected, result)

	result2 := ll.NeighbourOnCount(0, 0)
	expected2 := 1
	require.Equal(t, expected2, result2)

	result3 := ll.NeighbourOnCount(2, 2)
	expected3 := 2
	require.Equal(t, expected3, result3)
}

func TestCount(t *testing.T) {
	input := [][]int{}

	row0 := []int{0, 1, 0}
	row1 := []int{0, 0, 1}
	row2 := []int{1, 1, 1}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)

	ll := lights.New(input)

	result := ll.CountLights()
	expected := 5
	require.Equal(t, expected, result)
}

func TestOverrideCorners(t *testing.T) {
	input := [][]int{}
	row0 := []int{0, 0, 0}
	row1 := []int{0, 0, 0}
	row2 := []int{0, 0, 0}

	input = append(input, row0)
	input = append(input, row1)
	input = append(input, row2)

	ll := lights.New(input)

	expected := [][]int{}
	_row0 := []int{1, 0, 1}
	_row1 := []int{0, 0, 0}
	_row2 := []int{1, 0, 1}

	expected = append(expected, _row0)
	expected = append(expected, _row1)
	expected = append(expected, _row2)

	result := ll.OverrideCorners()
	require.Equal(t, expected, result)
}
