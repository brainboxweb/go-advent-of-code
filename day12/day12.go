package day12

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

func Count(input string) int {
	re := regexp.MustCompile("-?[0-9]+")
	matches := re.FindAllString(input, -1)
	var total int
	for _, value := range matches {
		integer, _ := strconv.Atoi(value)
		total += integer
	}
	return total
}

func Count2(input string) int {
	var f []any
	err := json.Unmarshal([]byte(input), &f)
	if err != nil {
		panic("Not expected")
	}
	total := parseArray(f)

	return total
}

func parseMap(aMap map[string]any) int {
	var total int
	for _, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]any:
			total += parseMap(val.(map[string]any))
		case []any:
			total += parseArray(val.([]any))
		default:
			if concreteVal == "red" {
				// just return - discard the subtotal for this entire node
				return 0
			}
			conc := fmt.Sprintf("%v", concreteVal)
			integer, err := strconv.Atoi(conc)
			if err != nil {
				continue
			}
			total += integer
		}
	}
	return total
}

func parseArray(anArray []any) int {
	var total int
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]any:
			total += parseMap(val.(map[string]any))
		case []any:
			total += parseArray(val.([]any))
		default:
			conc := fmt.Sprintf("%v", concreteVal)
			integer, err := strconv.Atoi(conc)
			if err != nil {
				continue
			}
			total += integer
		}
	}
	return total
}
