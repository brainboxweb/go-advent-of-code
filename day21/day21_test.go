package day21_test

import (
	"testing"

	"github.com/brainboxweb/advent/day21"
	"github.com/stretchr/testify/require"
)

func TestPlayToWin(t *testing.T) { // <-- Part 1
	cost := day21.PlayToWin()
	expected := 111
	require.Equal(t, expected, cost)
}

func TestPlayToLose(t *testing.T) { // <-- Part 2
	cost := day21.PlayToLose()
	expected := 188
	require.Equal(t, expected, cost)
}
