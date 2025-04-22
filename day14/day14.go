package day14

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/brainboxweb/advent/day14/race"
)

func Part1(input string, time int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	reindeers := []race.Reindeer{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)
		name := parsed[0]
		speed, _ := strconv.Atoi(parsed[1])
		flyTime, _ := strconv.Atoi(parsed[2])
		restTime, _ := strconv.Atoi(parsed[3])

		reindeer := race.NewReindeer(name, speed, flyTime, restTime)
		reindeers = append(reindeers, *reindeer)
	}

	maxDistance := 0
	for _, reindeer := range reindeers {
		distance := reindeer.GetDistance(time)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}

func Part2(input string, time int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	r := race.Race{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)
		name := parsed[0]
		speed, _ := strconv.Atoi(parsed[1])
		flyTime, _ := strconv.Atoi(parsed[2])
		restTime, _ := strconv.Atoi(parsed[3])
		reindeer := race.NewReindeer(name, speed, flyTime, restTime)
		r.AddReindeer(reindeer)
	}
	r.RunRace(time)
	return r.GetTopScore()
}

func parse(phrase string) []string {
	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")

	return []string{tokens[0], tokens[3], tokens[6], tokens[13]}
}
