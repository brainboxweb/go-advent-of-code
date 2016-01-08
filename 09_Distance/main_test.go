package main

import (
	"reflect"
	"testing"
)

var test = []struct {
	input    string
	expected []string
}{
	{
		"AlphaCentauri to Snowdin = 66",
		[]string{"AlphaCentauri", "Snowdin", "66"},
	},
}

func TestParse(t *testing.T) {
	for _, test := range test {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

var test2 = []struct {
	input    string
	expected int
}{
	{
		"AlphaCentauri to Snowdin = 66\nSnowdin to Tambi = 22\nTambi to Faerun = 39",
		22,
	},
	{
		day9data,
		141,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		if actual := Run(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day9data = `AlphaCentauri to Snowdin = 66
AlphaCentauri to Tambi = 28
AlphaCentauri to Faerun = 60
AlphaCentauri to Norrath = 34
AlphaCentauri to Straylight = 34
AlphaCentauri to Tristram = 3
AlphaCentauri to Arbre = 108
Snowdin to Tambi = 22
Snowdin to Faerun = 12
Snowdin to Norrath = 91
Snowdin to Straylight = 121
Snowdin to Tristram = 111
Snowdin to Arbre = 71
Tambi to Faerun = 39
Tambi to Norrath = 113
Tambi to Straylight = 130
Tambi to Tristram = 35
Tambi to Arbre = 40
Faerun to Norrath = 63
Faerun to Straylight = 21
Faerun to Tristram = 57
Faerun to Arbre = 83
Norrath to Straylight = 9
Norrath to Tristram = 50
Norrath to Arbre = 60
Straylight to Tristram = 27
Straylight to Arbre = 81
Tristram to Arbre = 90`
