package day23

import (
	"github.com/brainboxweb/advent/day23/turin"
)

func Run(input string, intialValueForA int) int {
	_, endValueForB := turin.Run(input, intialValueForA)

	return endValueForB
}
