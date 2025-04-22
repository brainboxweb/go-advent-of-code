package day15_test

import (
	"testing"

	"github.com/brainboxweb/advent/day15"
)

func TestRun(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
			62842880,
		},
		{
			day15data,
			21367368,
		},
	}
	for _, test := range test {
		if actual := day15.Part1(test.input, 0); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

func TestRun500(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
			57600000,
		},
		{
			day15data,
			1766400,
		},
	}
	for _, test := range test {
		if actual := day15.Part1(test.input, 500); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

const day15data = `Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
Butterscotch: capacity 0, durability 5, flavor -3, texture 0, calories 3
Chocolate: capacity 0, durability 0, flavor 5, texture -1, calories 8
Candy: capacity 0, durability -1, flavor 0, texture 5, calories 8`
