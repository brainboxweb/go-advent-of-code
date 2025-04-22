package day14

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	/*
		Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
		Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
	*/
	var test = []struct {
		input    string
		expected []string
	}{
		{
			"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
			[]string{"Comet", "14", "10", "127"},
		},
		{
			"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			[]string{"Dancer", "16", "11", "162"},
		},
	}

	for _, test := range test {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}
