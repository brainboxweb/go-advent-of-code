package main

import (
	"fmt"
)

func Deliver(input string) int {
	x := 0
	y := 0
	m := make(map[string]bool)
	key := fmt.Sprintf("%d,%d", x, y)
	m[key] = true
	for _, r := range input {
		symbol := string(r)
		switch symbol {
		case ">":
			x++
		case "^":
			y++
		case "<":
			x--
		case "v":
			y--
		}
		key := fmt.Sprintf("%d,%d", x, y)
		m[key] = true
	}
	return len(m)
}

var visited = make(map[string]bool)

func RoboDeliver(input string) int {
	for k := range visited {
		delete(visited, k)
	}
	visited["0,0"] = true
	santa := Deliverer{0, 0, "santa"}
	robot := Deliverer{0, 0, "robot"}
	counter := 0
	for _, symbol := range input {
		counter++
		if counter%2 == 0 {
			santa.Move(string(symbol))
		} else {
			robot.Move(string(symbol))
		}
	}
	return len(visited)
}

type Deliverer struct {
	X    int
	Y    int
	Name string
}

func (d *Deliverer) Move(symbol string) {
	switch symbol {
	case ">":
		d.X++
	case "^":
		d.Y++
	case "<":
		d.X--
	case "v":
		d.Y--
	}
	key := fmt.Sprintf("%d,%d", d.X, d.Y)
	visited[key] = true
}
