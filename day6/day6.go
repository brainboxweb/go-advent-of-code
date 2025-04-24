package day6

import (
	"github.com/brainboxweb/advent/day6/lights"
)

func Part1(input []string) int {
	ll := lights.New(input)
	return ll.GetLightCount()
}

func Part2(input []string) int {
	ll := lights.NewAdvanced(input)
	return ll.GetTotalBrightness()
}
