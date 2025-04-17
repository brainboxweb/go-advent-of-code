package day13

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/brainboxweb/advent/day13/dinner"
)

func Run(input string) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	relationships := dinner.Relationships{}
	persons := dinner.Persons{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)

		persons.AddPerson(parsed[0])
		persons.AddPerson(parsed[1])

		happiness, _ := strconv.Atoi(parsed[2])
		relationships.AddRelationship(parsed[0], parsed[1], happiness)
	}
	return relationships.GetHappiest(persons)
}

func RunWithGuest(input string) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	relationships := dinner.Relationships{}
	persons := dinner.Persons{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)

		persons.AddPerson(parsed[0])
		persons.AddPerson(parsed[1])

		happiness, _ := strconv.Atoi(parsed[2])
		relationships.AddRelationship(parsed[0], parsed[1], happiness)

		// extraGuest
		persons.AddPerson("Guest")

	}
	return relationships.GetHappiest(persons)
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
