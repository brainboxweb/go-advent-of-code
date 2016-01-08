package main

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	Action string
	Start  string
	End    string
}

func (i *instruction) coordsStart() (x, y int) {
	point := strings.Split(i.Start, ",")
	x, _ = strconv.Atoi(point[0])
	y, _ = strconv.Atoi(point[1])
	return x, y
}

func (i *instruction) coordsEnd() (x, y int) {
	point := strings.Split(i.End, ",")
	x, _ = strconv.Atoi(point[0])
	y, _ = strconv.Atoi(point[1])
	return x, y
}

type light struct {
	Level int
}


func GetLightCount(input string) int {

	var lights = make(map[string]light)

	b := bytes.NewBufferString(input)

	adjustedInstructions := false
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		processInstruction(lights, adjustedInstructions, scanner.Text())
	}

	return len(lights)
}

func GetTotalBrightness(input string) int {

	var lights = make(map[string]light)

	b := bytes.NewBufferString(input)

	scanner := bufio.NewScanner(b)

	adjustedInstructions := true
	for scanner.Scan() {
		processInstruction(lights, adjustedInstructions, scanner.Text())
	}

	brightnessCount := 0
	for _, light := range lights {
		brightnessCount += light.Level
	}
	return brightnessCount
}



func processInstruction(lights map[string]light, adjustedInstructions bool, input string) {
	instruction := parseInstruction(input)
	x1, y1 := instruction.coordsStart()
	x2, y2 := instruction.coordsEnd()
	for xx := x1; xx <= x2; xx++ {
		for yy := y1; yy <= y2; yy++ {
			key := fmt.Sprintf("%d,%d", xx, yy)
			processlight(lights, key, instruction.Action, adjustedInstructions)
		}
	}
}

func processlight(lights map[string]light, key, instruction string, adjustedInstructions bool) {

	switch adjustedInstructions{
	case false:
		switch instruction {
		case "on":
			switchOn(lights, key)
		case "off":
			switchOff(lights, key)
		case "toggle":
			toggle(lights, key)
		}
	case true:
		switch instruction {
		case "on":
			switchOnAdjusted(lights, key)
		case "off":
			switchOffAdjusted(lights, key)
		case "toggle":
			toggleAdjusted(lights, key)
		}
	}
}

func switchOn(lights map[string]light, key string) {
	//if exisits, do nothing
	if _, ok := lights[key]; ok {
		return
	}
	//otherwise, create
	lights[key] = light{}
}

func switchOff(lights map[string]light, key string) {
	if _, ok := lights[key]; ok {
		delete(lights, key)
	}
}

func toggle(lights map[string]light, key string) {
	//If exists... delete it
	if _, ok := lights[key]; ok {
		delete(lights, key)
		return
	}
	//otherwise, create it
	lights[key] = light{}
}

//turn off 370,39 through 425,839
//turn off 464,858 through 833,915
func parseInstruction(phrase string) instruction {
	r, _ := regexp.Compile(`(\d{1,3},\d{1,3})`)
	result := r.FindAllString(phrase, -1)
	action := ""
	switch {
	case strings.Contains(phrase, "turn on"):
		action = "on"
	case strings.Contains(phrase, "turn off"):
		action = "off"
	default:
		action = "toggle"
	}
	instruction := instruction{Action: action, Start: result[0], End: result[1]}
	return instruction
}


//--------- PART B


func switchOnAdjusted(lights map[string]light, key string) {
//	//if exists, increase level
	if lightCopy, ok := lights[key]; ok {
		lightCopy.Level++
		lights[key] = lightCopy
	} else {
		//otherwise, create
		newLight := light{Level: 1}
		lights[key] = newLight
	}
}


func switchOffAdjusted(lights map[string]light, key string) {
	//if exists, reduce level
	if light, ok := lights[key]; ok {
		light.Level--
		if light.Level < 1 {
			delete(lights, key)
			return
		}
		lights[key] = light
	}
}


func toggleAdjusted(lights map[string]light, key string) {
	//increase twice
	switchOnAdjusted(lights, key)
	switchOnAdjusted(lights, key)
}
