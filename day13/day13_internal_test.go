package day13

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
			"Alice would gain 54 happiness units by sitting next to Bob.",
			[]string{"Alice", "Bob", "54"},
		},
		{
			"Alice would lose 79 happiness units by sitting next to Carol.",
			[]string{"Alice", "Carol", "-79"},
		},
	}
	for _, test := range test {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}
