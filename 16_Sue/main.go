package main

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
	"strings"
	//	"fmt"
)

func run(input string, known string, advanced bool) int {

	//The Sues
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	sues := []Sue{}

	for scanner.Scan() {
		text := scanner.Text()
		sue := parse(text)

		sues = append(sues, sue)
	}

	// The know properties
	b2 := bytes.NewBufferString(known)
	scanner2 := bufio.NewScanner(b2)

	knowns := []map[string]int{}
	for scanner2.Scan() {
		text := scanner2.Text()
		known := parseKnown(text)

		knowns = append(knowns, known)
	}

	return findMatch(sues, knowns, advanced)
}

func findMatch(sues []Sue, knowns []map[string]int, advanced bool) int {

	type match struct {
		sue        Sue
		matchCount int
	}

	matches := map[int]Sue{}

	for _, sue := range sues {
		matchCount := 0
		for _, known := range knowns {
			//Range over the knowns
			for key, value := range known {
				//range over the sue properties
				for key2, value2 := range sue.properies {

					//For part towo, needtochnange the match types
					/*
						In particular, the cats and trees readings indicates that there are greater than that many
						(due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians
						 and goldfish readings indicate that there are fewer than that many (due to the modial
						 interaction of magnetoreluctance).
					*/

					if key != key2 {
						continue
					}

					if advanced == false {
						if value == value2 {
							matchCount++
						}
					} else {
						switch key {
						case "cats":
							fallthrough
						case "trees":
							//greater
							if value2 > value {
								matchCount++
							}
						case "pomeranians":
							fallthrough
						case "goldfish":
							//less
							if value2 < value {
								matchCount++
							}
						default:
							if value == value2 {
								matchCount++
							}
						}
					}
				}
			}
		}
		if matchCount == 0 {
			continue
		}
		//Store the matches
		matches[matchCount] = sue
	}

	//Sort
	var keys []int
	for k, _ := range matches {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//Need the last one = the nest match
	requiredSue := matches[keys[len(keys)-1]]

	return requiredSue.index
}

//Sue 1: goldfish: 9, cars: 0, samoyeds: 9
func parse(phrase string) Sue {
	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")

	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.Trim(tokens[i], ":")
		tokens[i] = strings.Trim(tokens[i], ",")
	}

	index, _ := strconv.Atoi(tokens[1])

	properties := make(map[string]int)

	val, _ := strconv.Atoi(tokens[3])
	properties[tokens[2]] = val

	val2, _ := strconv.Atoi(tokens[5])
	properties[tokens[4]] = val2

	val3, _ := strconv.Atoi(tokens[7])
	properties[tokens[6]] = val3

	sue := Sue{index, properties}

	return sue
}

// children: 3
// cats: 7
func parseKnown(phrase string) map[string]int {

	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")

	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.Trim(tokens[i], ":")
	}

	known := make(map[string]int)

	value, _ := strconv.Atoi(tokens[1])
	known[tokens[0]] = value

	return known
}

//Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
type Sue struct {
	index     int
	properies map[string]int
}
