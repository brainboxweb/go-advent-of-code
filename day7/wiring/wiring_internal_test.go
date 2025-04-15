package wiring

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{
			"123 -> a",
			[]string{"123", "->", "a"},
		},
		{
			"nx LSHIFT 1 -> a",
			[]string{"nx", "LSHIFT", "1", "->", "a"},
		},
		{
			"NOT im -> in",
			[]string{"NOT", "im", "->", "in"},
		},
		{
			"x AND y -> d",
			[]string{"x", "AND", "y", "->", "d"},
		},
		{
			"x OR y -> d",
			[]string{"x", "OR", "y", "->", "d"},
		},
	}
	for _, test := range tests {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

func TestGetOperator(t *testing.T) {
	var test = []struct {
		input    []string
		expected string
	}{
		{
			[]string{"123", "->", "a"},
			"ASSIGN",
		},
		{
			[]string{"nx", "LSHIFT", "1", "->", "a"},
			"LSHIFT",
		},
		{
			[]string{"nx", "RSHIFT", "1", "->", "a"},
			"RSHIFT",
		},
		{
			[]string{"NOT", "im", "->", "in"},
			"NOT",
		},
		{
			[]string{"x", "AND", "y", "->", "d"},
			"AND",
		},
	}
	for _, test := range test {
		if actual := getOperator(test.input); actual != test.expected {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

