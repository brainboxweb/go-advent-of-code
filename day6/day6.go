package day6

import (
	"github.com/brainboxweb/advent/day6/lights"
)

func Part1(input []string) int {
	lights := lights.New(input)
	return lights.GetLightCount()
}

func Part2(input []string) int {
	lights := lights.NewAdvanced(input)
	return lights.GetTotalBrightness()
}
