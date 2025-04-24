package lights

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func New(data []string) *Lights {
	var ls = make(map[string]light)
	lights := Lights{}
	lights.lights = ls
	for _, line := range data {
		lights.processInstruction(line)
	}
	return &lights
}

func NewAdvanced(data []string) *Lights {
	var ls = make(map[string]light)
	lights := Lights{isAdvanced: true}
	lights.lights = ls
	for _, line := range data {
		lights.processInstruction(line)
	}
	return &lights
}

type Lights struct {
	isAdvanced bool
	lights     map[string]light
}

type light struct {
	Level int
}

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

func (ll *Lights) GetLightCount() int {
	return len(ll.lights)
}

func (ll *Lights) GetTotalBrightness() int {
	if !ll.isAdvanced {
		panic("not supported")
	}
	brightnessCount := 0
	for _, light := range ll.lights {
		brightnessCount += light.Level
	}
	return brightnessCount
}

func (ll *Lights) processInstruction(input string) {
	instruction := parseInstruction(input)
	x1, y1 := instruction.coordsStart()
	x2, y2 := instruction.coordsEnd()
	for xx := x1; xx <= x2; xx++ {
		for yy := y1; yy <= y2; yy++ {
			key := fmt.Sprintf("%d,%d", xx, yy)
			ll.processlight(key, instruction.Action)
		}
	}
}

func (ll *Lights) processlight(key, instruction string) {
	switch instruction {
	case "on":
		ll.switchOn(key)
	case "off":
		ll.switchOff(key)
	case "toggle":
		ll.toggle(key)
	}
}

func (ll *Lights) switchOn(key string) {
	if ll.isAdvanced {
		// if exists, increase level
		_, ok := ll.lights[key]
		if ok {
			light := ll.lights[key]
			light.Level++
			ll.lights[key] = light
			return
		}

		light := light{Level: 1}
		ll.lights[key] = light
		return
	}
	// if exisits, do nothing
	if _, ok := ll.lights[key]; ok {
		return
	}
	// otherwise, create
	ll.lights[key] = light{}
}

func (ll *Lights) switchOff(key string) {
	if ll.isAdvanced {
		// if exists, reduce level
		if light, ok := ll.lights[key]; ok {
			light.Level--
			if light.Level < 1 {
				delete(ll.lights, key)
				return
			}
			ll.lights[key] = light
		}
		return
	}

	delete(ll.lights, key)
}

func (ll *Lights) toggle(key string) {
	if ll.isAdvanced {
		// increase twice
		ll.switchOn(key)
		ll.switchOn(key)
		return
	}
	// If exists... delete it
	if _, ok := ll.lights[key]; ok {
		delete(ll.lights, key)
		return
	}
	// otherwise, create it
	ll.lights[key] = light{}
}

func parseInstruction(phrase string) instruction {
	r := regexp.MustCompile(`(\d{1,3},\d{1,3})`)
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
