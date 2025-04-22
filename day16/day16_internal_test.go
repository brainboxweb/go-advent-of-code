package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1

Sue 1: goldfish: 9, cars: 0, samoyeds: 9
Sue 2: perfumes: 5, trees: 8, goldfish: 8
*/

func TestParse(t *testing.T) {
	sue1 := `Sue 1: goldfish: 9, cars: 0, samoyeds: 9`
	params := make(map[string]int)
	params["cars"] = 0
	params["goldfish"] = 9
	params["samoyeds"] = 9

	expected := Sue{1, params}
	result := parse(sue1)

	require.Equal(t, expected, result)
}
