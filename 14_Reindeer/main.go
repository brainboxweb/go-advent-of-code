package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	//	"fmt"
	"sort"
)

func run(input string, time int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	reindeers := []Reindeer{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)
		name := parsed[0]
		speed, _ := strconv.Atoi(parsed[1])
		flyTime, _ := strconv.Atoi(parsed[2])
		restTime, _ := strconv.Atoi(parsed[3])
		reindeer := Reindeer{name: name, speed: speed, flyTime: flyTime, restTime: restTime}
		reindeers = append(reindeers, reindeer)
	}

	maxDistance := 0
	for _, reindeer := range reindeers {
		distance := reindeer.getDistance(time)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}

func run2(input string, time int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	race := Race{}
	for scanner.Scan() {
		text := scanner.Text()
		parsed := parse(text)
		name := parsed[0]
		speed, _ := strconv.Atoi(parsed[1])
		flyTime, _ := strconv.Atoi(parsed[2])
		restTime, _ := strconv.Atoi(parsed[3])
		reindeer := Reindeer{name: name, speed: speed, flyTime: flyTime, restTime: restTime}
		race.addReindeer(&reindeer)
	}
	race.runRace(time)
	return race.getTopScore()
}

func parse(phrase string) []string {
	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")

	return []string{tokens[0], tokens[3], tokens[6], tokens[13]}
}

type Reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
}

func (r *Reindeer) getDistance(time int) int {

	totalFlyTime := 0

	//Loops
	loopTime := r.flyTime + r.restTime
	chunks := time / loopTime // Taking advantance of integer maths
	totalFlyTime += chunks * r.flyTime

	//Partials
	remainderTime := time - (chunks * loopTime)
	if remainderTime > r.flyTime {
		totalFlyTime += r.flyTime
	} else {
		totalFlyTime += remainderTime
	}

	return totalFlyTime * r.speed
}

type Race struct {
	scores map[*Reindeer]int
}

func (race *Race) addReindeer(reindeer *Reindeer) {

	//Avoid nil map panic
	if race.scores == nil {
		race.scores = make(map[*Reindeer]int)
	}
	race.scores[reindeer] = 0
}

func (race *Race) runRace(time int) {

	//Award points every second
	for i := 1; i < time+1; i++ {
		dists := distances{}

		for rein, _ := range race.scores {
			dist := distance{rein, rein.getDistance(i)}
			dists = append(dists, dist)
		}

		//Dists still hse referecen to the real raindeer
		sort.Sort(dists)

		winningDistance := dists[0].Distance

		for i := 0; i < len(dists); i++ {
			if winningDistance == dists[i].Distance {
				theReindeer := dists[i].Reindeer
				race.scores[theReindeer]++
			}

		}
	}
}

func (race *Race) getTopScore() int {
	topScore := 0
	for _, score := range race.scores {
		if score > topScore {
			topScore = score
		}
	}
	return topScore
}

//For sorting
type distance struct {
	Reindeer *Reindeer
	Distance int
}

type distances []distance

func (d distances) Len() int      { return len(d) }
func (d distances) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d distances) Less(i, j int) bool {
	if d[i].Distance > d[j].Distance { //reverse search
		return true
	} else {
		return false
	}
}
