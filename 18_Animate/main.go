package main

import (
	"bufio"
	"bytes"
	//	"fmt"
	"strings"
)

func run(input string, loops int, override bool) int {

	lights := parse(input)

	if override == true {
		lights = overrideCorners(lights)
	}

	for i := 0; i < loops; i++ {
		lights = switchLights(lights)
		if override == true {
			lights = overrideCorners(lights)
		}
	}
	return countLights(lights)
}

func countLights(lights [][]int) int {

	total := 0
	for _, row := range lights {

		for _, item := range row {
			if item == 1 {
				total++
			}
		}
	}

	return total
}

/*
.#.#.#
...##.
#....#
..#...
#.#..#
####..
*/
func parse(input string) [][]int {

	coords := [][]int{}

	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	for scanner.Scan() {
		line := scanner.Text()

		lights := strings.Split(line, "")
		row := []int{}
		for _, light := range lights {

			if light == "." {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}

		coords = append(coords, row)
	}

	return coords

}

func switchLights(input [][]int) [][]int {

	output := deepCopy(input)

	for y := 0; y < len(input); y++ {

		for x := 0; x < len(input[y]); x++ {

			neightbourOnCount := neighbourOnCount(input, x, y)

			switch input[y][x] {
			case 1: //bulb in on
				if neightbourOnCount == 2 || neightbourOnCount == 3 {
					//no change
				} else {

					output[y][x] = 0
				}
			case 0: //bulb off
				if neightbourOnCount == 3 {
					output[y][x] = 1
				}
			default:
				panic("really??!?!?!")
			}
		}
	}

	return output
}

func overrideCorners(input [][]int) [][]int {

	output := deepCopy(input)

	ymax := len(input) - 1
	xmax := len(input[0]) - 1

	output[0][0] = 1
	output[ymax][0] = 1
	output[0][xmax] = 1
	output[ymax][xmax] = 1

	return output
}

func deepCopy(input [][]int) [][]int {

	output := [][]int{}

	for _, row := range input {
		row2 := make([]int, len(row))
		copy(row2, row)
		output = append(output, row2)
	}
	return output
}

func neighbourOnCount(input [][]int, x, y int) int {

	onCount := 0

	ymax := len(input) - 1
	xmax := len(input[0]) - 1

	for yy := y - 1; yy < y+2; yy++ {

		for xx := x - 1; xx < x+2; xx++ {
			if xx == x && yy == y { //ignore current
				continue
			}
			if xx < 0 || yy < 0 || xx > xmax || yy > ymax { //edges
				continue
			}
			if input[yy][xx] == 1 {
				onCount++
			}
		}
	}

	return onCount
}
