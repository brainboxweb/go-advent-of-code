package day2

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
	"strings"
)

func Wrapping(input string) (int, int) {

	b := bytes.NewBufferString(input)
	totalArea := 0
	totalRibbon := 0

	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		theText := scanner.Text()
		items := strings.Split(theText, "x")
		a, _ := strconv.Atoi(items[0])
		b, _ := strconv.Atoi(items[1])
		c, _ := strconv.Atoi(items[2])
		area := 2 * (a*b + a*c + b*c)
		sides := []int{a, b, c}
		sort.Ints(sides)
		extra := sides[0] * sides[1]
		totalArea += area + extra

		//Ribbon
		ribbon := 2 * (sides[0] + sides[1])
		bow := a * b * c
		totalRibbon += ribbon + bow
	}
	return totalArea, totalRibbon
}
