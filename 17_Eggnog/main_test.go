package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to
move it into smaller containers. You take an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there
are four ways to do it:

15 and 10
20 and 5 (the first 5)
20 and 5 (the second 5)
15, 5, and 5
Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of eggnog?
*/

var test2 = []struct {
	input            string
	volume           int
	expected         int
	expectedMinCount int
}{
	{
		`10
5`,
		15,
		1,
		1,
	},
	{
		`20
15
10
5
5`,
		25,
		4,
		3,
	},
	{
		day17data,
		150,
		654,
		57,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		if actual, _ := run(test.input, test.volume); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
		if _, actual := run(test.input, test.volume); actual != test.expectedMinCount {
			t.Errorf("Min count(%s) = %d, expected %d.",
				test.input, actual, test.expectedMinCount)
		}
	}
}

//The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to
//move it into smaller containers. You take an inventory of the capacities of the available containers.
//
//For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there
//are four ways to do it:
//15 and 10
//20 and 5 (the first 5)
//20 and 5 (the second 5)
//15, 5, and 5

func TestRecurse(t *testing.T) {

	bottle := Bottle{"a", 20}
	bottle2 := Bottle{"b", 5}
	bottle3 := Bottle{"c", 10}
	bottle4 := Bottle{"d", 5}
	bottle5 := Bottle{"e", 15}

	bottles := []Bottle{}
	bottles = append(bottles, bottle)
	bottles = append(bottles, bottle2)
	bottles = append(bottles, bottle3)
	bottles = append(bottles, bottle4)
	bottles = append(bottles, bottle5)

	breadcrumb := ""

	volume := 25

	matches := []string{}
	result := recurse(breadcrumb, bottles, volume, matches)

	expected := []string{}
	expected = append(expected, "ab")
	expected = append(expected, "ad")
	expected = append(expected, "ba")
	expected = append(expected, "bde")
	expected = append(expected, "bed")
	expected = append(expected, "ce")
	expected = append(expected, "da")
	expected = append(expected, "dbe")
	expected = append(expected, "deb")
	expected = append(expected, "ebd")
	expected = append(expected, "ec")
	expected = append(expected, "edb")

	require.Equal(t, expected, result)
}

var test3 = []struct {
	input    []string
	expected int
	minCount int
}{
	{
		[]string{"abc", "cab", "bac"},
		1,
		1,
	},
	{
		[]string{"abc", "cab", "bac", "ab"},
		2,
		1,
	},
}

func TestDeDupe(t *testing.T) {
	for _, test := range test3 {
		if actual, _ := dedupe(test.input); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
		if _, actual := dedupe(test.input); actual != test.minCount {
			t.Errorf("Dedupe minCount (%s) = %d, expected %d.",
				test.input, actual, test.minCount)
		}
	}
}

const day17data = `50
44
11
49
42
46
18
32
26
40
21
7
18
43
10
47
36
24
22
40`
