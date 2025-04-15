package lights_test

import (
	"testing"

	"github.com/brainboxweb/advent/day6/lights"
)

func TestLights(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{"turn on 0,0 through 1,1"},
			4,
		},
		{
			[]string{
				"turn on 0,0 through 1,1",
				"turn off 1,1 through 1,1",
			},
			3,
		},
		{
			[]string{"turn on 0,0 through 2,2"},
			9,
		},
		{
			[]string{
				"turn on 0,0 through 2,2",
				"toggle 0,0 through 1,1",
			},
			5,
		},
		{
			[]string{"turn on 0,0 through 999,999"},
			1000000,
		},
		{
			[]string{"toggle 0,0 through 999,0"},
			1000,
		},
		{
			[]string{
				"turn on 0,0 through 999,999",
				"toggle 0,0 through 999,0",
			},
			999000,
		},
	}
	for _, test := range tests {
		lights := lights.New(test.input)
		if actual := lights.GetLightCount(); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

func TestAdvancedLights(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{"turn on 0,0 through 0,0"},
			1,
		},
		{
			[]string{"toggle 0,0 through 999,999"},
			2000000,
		},
	}
	for _, test := range tests {
		lights := lights.NewAdvanced(test.input)
		if actual := lights.GetTotalBrightness(); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}
