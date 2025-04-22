package day14_test

import (
	"testing"

	"github.com/brainboxweb/advent/day14"
)

func TestPart1(t *testing.T) {
	var test = []struct {
		input    string
		time     int
		expected int
	}{
		{
			`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
			1000,
			1120,
		},
		{
			day14data,
			2503,
			2655,
		},
	}
	for _, test := range test {
		if actual := day14.Part1(test.input, test.time); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	var test = []struct {
		input    string
		time     int
		expected int
	}{
		{
			`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
			1000,
			689,
		},
		{
			day14data,
			2503,
			1059,
		},
	}
	for _, test := range test {
		if actual := day14.Part2(test.input, test.time); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day14data = `Vixen can fly 8 km/s for 8 seconds, but then must rest for 53 seconds.
Blitzen can fly 13 km/s for 4 seconds, but then must rest for 49 seconds.
Rudolph can fly 20 km/s for 7 seconds, but then must rest for 132 seconds.
Cupid can fly 12 km/s for 4 seconds, but then must rest for 43 seconds.
Donner can fly 9 km/s for 5 seconds, but then must rest for 38 seconds.
Dasher can fly 10 km/s for 4 seconds, but then must rest for 37 seconds.
Comet can fly 3 km/s for 37 seconds, but then must rest for 76 seconds.
Prancer can fly 9 km/s for 12 seconds, but then must rest for 97 seconds.
Dancer can fly 37 km/s for 1 seconds, but then must rest for 36 seconds.`
