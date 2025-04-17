package day9

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var test = []struct {
		input    string
		expected []string
	}{
		{
			"AlphaCentauri to Snowdin = 66",
			[]string{"AlphaCentauri", "Snowdin", "66"},
		},
	}
	for _, test := range test {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}
