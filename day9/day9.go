package day9

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/nightlyone/permutation"
)

func Shortest(input string) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	routes := Routes{}
	locations := Locations{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)

		locations.AddLocation(parsed[0])
		locations.AddLocation(parsed[1])

		distance, _ := strconv.Atoi(parsed[2])

		routes.AddRoute(parsed[0], parsed[1], distance)
	}

	//Get locations in a map
	locationsMap := make(map[int]string)
	counter := 0
	for _, locn := range locations.Locations {
		locationsMap[counter] = locn
		counter++
	}

	//create a n int list
	keys := intSlice{}
	for i := 0; i < len(locationsMap); i++ {
		keys = append(keys, i)
	}

	// Permutation of simple slice
	perm := permutation.New(keys)
	shortest := 1000000

	// iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {

		indexes := seqToSlice(seq)

		newLocations := []string{}
		for _, index := range indexes {
			newLocations = append(newLocations, locations.Locations[index])
		}

		distance := getDistance(newLocations, routes)
		if distance == 0 {
			continue
		}

		if distance < shortest {
			shortest = distance
		}
	}
	return shortest
}

func Longest(input string) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	routes := Routes{}
	locations := Locations{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)

		locations.AddLocation(parsed[0])
		locations.AddLocation(parsed[1])

		distance, _ := strconv.Atoi(parsed[2])

		routes.AddRoute(parsed[0], parsed[1], distance)
	}

	//Get locations in a map
	locationsMap := make(map[int]string)
	counter := 0
	for _, locn := range locations.Locations {
		locationsMap[counter] = locn
		counter++
	}

	//create a n int list
	keys := intSlice{}
	for i := 0; i < len(locationsMap); i++ {
		keys = append(keys, i)
	}

	// Permutation of simple slice
	perm := permutation.New(keys)
	longest := 0

	// iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {

		indexes := seqToSlice(seq)

		newLocations := []string{}
		for _, index := range indexes {
			newLocations = append(newLocations, locations.Locations[index])
		}

		distance := getDistance(newLocations, routes)
		if distance == 0 {
			continue
		}

		if distance > longest {
			longest = distance
		}
	}
	return longest
}

func seqToSlice(seq permutation.Sequence) []int {

	hack := fmt.Sprint(seq)
	hack = strings.Trim(hack, "[]")
	indexes := strings.Split(hack, " ")

	indexesInt := []int{}

	for _, val := range indexes {

		indexInt, _ := strconv.Atoi(val)
		indexesInt = append(indexesInt, indexInt)
	}

	return indexesInt
}

func getDistance(locationsSlice []string, routes Routes) int {

	length := len(locationsSlice)
	distance := 0
	for i := 0; i < length-1; i++ {
		distance += routes.GetDistance(locationsSlice[i], locationsSlice[i+1])
	}

	return distance
}

func parse(phrase string) []string {
	tokens := strings.Split(phrase, " ")
	return []string{tokens[0], tokens[2], tokens[4]}
}

type Locations struct {
	Locations map[int]string
}

func (l *Locations) AddLocation(location string) {
	if l.Locations == nil {
		l.Locations = make(map[int]string)
	}

	//Test for dupes
	for _, locn := range l.Locations {
		if locn == location {
			return
		}
	}

	//Add it
	l.Locations[len(l.Locations)] = location
}

type Routes struct {
	routes map[string]int
}

func (r *Routes) AddRoute(one, two string, distance int) {

	if r.routes == nil {
		r.routes = make(map[string]int)
	}

	r.routes[one+","+two] = distance
	r.routes[two+","+one] = distance
}

func (r *Routes) GetDistance(one, two string) int {

	return r.routes[one+","+two]
}

///Permutations

// define custom type
type intSlice []int

// Implement the three functions from sort.Interface (part of permutation.Sequence interface)
func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Implement the remaining portions of permutation.Sequence interface
func (p intSlice) Equal(q permutation.Sequence) bool { return reflect.DeepEqual(p, q) }
func (p intSlice) Copy() permutation.Sequence {
	q := make(intSlice, len(p), len(p))
	copy(q, p)
	return q
}
