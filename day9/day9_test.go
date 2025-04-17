package day9_test

import (
	"testing"

	"github.com/brainboxweb/advent/day9"
)

func TestShortest(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			"AlphaCentauri to Snowdin = 66\nSnowdin to Tambi = 22\nTambi to Faerun = 39",
			22,
		},
		{
			day9data,
			117, // <-- Part 1
		},
	}
	for _, test := range test {
		if actual := day9.Shortest(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestLongest(t *testing.T) {
	var test = []struct {
		input    string
		expected int
	}{
		{
			day9data,
			909, // <-- Part 2
		},
	}
	for _, test := range test {
		if actual := day9.Longest(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day9data = `Faerun to Tristram = 65
Faerun to Tambi = 129
Faerun to Norrath = 144
Faerun to Snowdin = 71
Faerun to Straylight = 137
Faerun to AlphaCentauri = 3
Faerun to Arbre = 149
Tristram to Tambi = 63
Tristram to Norrath = 4
Tristram to Snowdin = 105
Tristram to Straylight = 125
Tristram to AlphaCentauri = 55
Tristram to Arbre = 14
Tambi to Norrath = 68
Tambi to Snowdin = 52
Tambi to Straylight = 65
Tambi to AlphaCentauri = 22
Tambi to Arbre = 143
Norrath to Snowdin = 8
Norrath to Straylight = 23
Norrath to AlphaCentauri = 136
Norrath to Arbre = 115
Snowdin to Straylight = 101
Snowdin to AlphaCentauri = 84
Snowdin to Arbre = 96
Straylight to AlphaCentauri = 107
Straylight to Arbre = 14
AlphaCentauri to Arbre = 46`
