package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

func Count(input string) int {
	re := regexp.MustCompile("-?[0-9]+") //-?[0-9]\d*(\.\d+)?
	matches := re.FindAllString(input, -1)
	var total int
	for _, value := range matches {
		integer, _ := strconv.Atoi(value)
		total += integer
	}
	return total
}

func Count2(input string) int {

	var f []interface{}
	err := json.Unmarshal([]byte(input), &f)
	if err != nil {
		panic("Didnt expect that")
	}
	total := parseArray(f) //Need to choose one or the other :(

	return total
}

func parseMap(aMap map[string]interface{}) int {
	var total int
	for _, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			total += parseMap(val.(map[string]interface{}))
		case []interface{}:
			total += parseArray(val.([]interface{}))
		default:

			if concreteVal == "red" {
				//just return - discard the subtotal for this entire node
				return 0
			}
			conc := fmt.Sprintf("%v", concreteVal)
			integer, err := strconv.Atoi(conc)
			if err != err {
				panic("oh dear...")
			}
			total += integer
		}
	}
	return total
}

func parseArray(anArray []interface{}) int {
	var total int
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			total += parseMap(val.(map[string]interface{}))
		case []interface{}:
			total += parseArray(val.([]interface{}))
		default:
			conc := fmt.Sprintf("%v", concreteVal)
			integer, err := strconv.Atoi(conc)
			if err != err {
				panic("oh dear...")
			}
			total += integer
		}
	}
	return total
}
