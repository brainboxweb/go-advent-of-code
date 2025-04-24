package main

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
	"strings"
)

func run(input string, volume int) (int, int) {

	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	labels := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	bottles := []Bottle{}
	i := 0
	for scanner.Scan() {
		if i > 24 {
			panic("That is going to be a problem")
		}
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		bottle := Bottle{labels[i], num}
		bottles = append(bottles, bottle)
		i++
	}

	//Let's recurse!
	breadcrumb := ""
	matches := []string{}
	results := recurse(breadcrumb, bottles, volume, matches)

	return dedupe(results)
}

type Matches struct {
	matches [][]Bottle
}

type Bottle struct {
	name string
	vol  int
}

// takes the current state of Bottle
// picks one...
// recurses through the others
func recurse(breadcrumb string, bottles []Bottle, volume int, matches []string) []string {

	for k, bottle := range bottles {

		newBreadcrumb := breadcrumb + bottle.name

		if bottle.vol > volume {
			//too big. move on
			continue
		}

		if bottle.vol == volume {
			//we have a winner. Add to current bottles
			matches = append(matches, newBreadcrumb)
			continue
		}

		//else... keep going!
		if len(bottles) > 0 {
			newVolume := volume - bottle.vol //must happend???
			//Copy bottles before recursing
			botts2 := make([]Bottle, len(bottles))
			copy(botts2, bottles)
			//Remove the current bottle
			botts2 = append(botts2[:k], botts2[k+1:]...)
			matches = recurse(newBreadcrumb, botts2, newVolume, matches)
		}
	}
	return matches
}

// Dedupe the string. Return the count
func dedupe(input []string) (int, int) {

	index := make(map[string]int)
	for _, string := range input {
		parts := strings.Split(string, "")
		sort.Strings(parts)
		index[strings.Join(parts, "")] = len(parts)
	}

	//Get all of the lengths
	lengths := []int{}
	for _, val := range index {
		lengths = append(lengths, val)
	}
	sort.Ints(lengths)
	minLength := lengths[0]
	//How many share this length?
	minCount := 0
	for _, val := range lengths {
		if val == minLength {
			minCount++
			continue
		}
		break //as soo as it changes... bail
	}

	return len(index), minCount
}
