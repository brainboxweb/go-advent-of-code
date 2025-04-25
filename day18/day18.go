package day18

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/brainboxweb/advent/day18/lights"
)

func Run(input string, loops int) int {
	data := parse(input)
	ll := lights.New(data)
	for i := 0; i < loops; i++ {
		ll.SwitchLights()
	}

	return ll.CountLights()
}

func RunWithOverride(input string, loops int) int {
	data := parse(input)
	ll := lights.New(data)
	ll.OverrideCorners()
	for i := 0; i < loops; i++ {
		ll.SwitchLights()
		ll.OverrideCorners()
	}

	return ll.CountLights()
}

func parse(input string) [][]int {
	coords := [][]int{}
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		line := scanner.Text()
		ll := strings.Split(line, "")
		row := []int{}
		for _, light := range ll {
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
