package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/nightlyone/permutation"
	"reflect"
	"strconv"
	"strings"
)

func Run(input string, extraGuest bool) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	relationships := Relationships{}
	persons := Persons{}
	for scanner.Scan() {

		text := scanner.Text()
		parsed := parse(text)

		persons.AddPerson(parsed[0])
		persons.AddPerson(parsed[1])

		happiness, _ := strconv.Atoi(parsed[2])
		relationships.addRelationship(parsed[0], parsed[1], happiness)

		if extraGuest == true {
			//Add an extra person
			persons.AddPerson("Guest")
		}
	}

	//Get persons in a map
	relationshipsMap := make(map[int]string)
	counter := 0
	for _, person := range persons.Persons {
		relationshipsMap[counter] = person
		counter++
	}

	//create a n int list
	keys := intSlice{}
	for i := 0; i < len(relationshipsMap); i++ {
		keys = append(keys, i)
	}

	// Permutation of simple slice
	perm := permutation.New(keys)

	happiest := 0
	// iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {

		indexes := seqToSlice(seq)
		newPersons := []string{}
		for _, index := range indexes {
			newPersons = append(newPersons, persons.Persons[index])
		}

		happiness := getHappiness(newPersons, relationships)
		if happiness > happiest {
			happiest = happiness
		}
	}
	return happiest
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

func getHappiness(personsSlice []string, relationships Relationships) int {

	length := len(personsSlice)
	happiness := 0
	for i := 0; i < length-1; i++ {
		happiness += relationships.GetHappiness(personsSlice[i], personsSlice[i+1])
		happiness += relationships.GetHappiness(personsSlice[i+1], personsSlice[i])
	}
	//Also need the happiness of the first and last
	happiness += relationships.GetHappiness(personsSlice[length-1], personsSlice[0])
	happiness += relationships.GetHappiness(personsSlice[0], personsSlice[length-1])

	return happiness
}

func parse(phrase string) []string {
	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")
	name1 := tokens[0]
	name2 := tokens[10]
	quant := tokens[3]
	if tokens[2] == "lose" {
		quant = "-" + quant
	}
	return []string{name1, name2, quant}
}

type Persons struct {
	Persons map[int]string
}

func (p *Persons) AddPerson(person string) {
	if p.Persons == nil {
		p.Persons = make(map[int]string)
	}
	//Test for dupes
	for _, locn := range p.Persons {
		if locn == person {
			return
		}
	}
	p.Persons[len(p.Persons)] = person
}

type Relationships struct {
	relationships map[string]int
}

func (r *Relationships) addRelationship(one, two string, happiness int) {
	//Avoid nil map panic
	if r.relationships == nil {
		r.relationships = make(map[string]int)
	}
	r.relationships[one+","+two] = happiness
}

func (r *Relationships) GetHappiness(one, two string) int {
	happiness := r.relationships[one+","+two]
	return happiness
}

//-------------------- Permutations ---------------------

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
