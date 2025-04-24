package main

import (
	"bufio"
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Parcel int

func (p Parcel) toString() string {
	return strconv.Itoa(int(p))
}

type Group struct {
	parcels   []Parcel
	maxWeight int
}

func (g Group) weight() (weight int) {
	for _, parcel := range g.parcels {
		weight += int(parcel)
	}
	return weight
}

func (g Group) len() int {
	return len(g.parcels)
}

func (g *Group) Add(p Parcel) bool {
	//Add... unless it exceeds maxweight
	if g.weight()+int(p) > g.maxWeight {
		return false
	}

	g.parcels = append(g.parcels, p)
	return true
}

type Sleigh struct {
	groups []Group
}

func (s *Sleigh) addParcels(parcels []Parcel) bool {
	//Need descending order???
	//tototal weufght:
	totalWeight := 0
	for _, parcel := range parcels {
		totalWeight += int(parcel)
	}
	groupWeight := totalWeight / len(s.groups)
	//Inform the Groups
	for k, _ := range s.groups {
		s.groups[k].maxWeight = groupWeight
	}

	for _, parcel := range parcels {
		for k, _ := range s.groups {
			ok := s.groups[k].Add(parcel)
			if ok {
				break //parcel placed successfully
			}
		}
	}

	//Check that ALL parcels have been placed.
	parcelCount := 0
	for _, group := range s.groups {
		parcelCount += len(group.parcels)
	}
	if parcelCount != len(parcels) {
		return false
	}

	//Check for balance
	for _, group := range s.groups {
		if groupWeight != group.weight() {
			return false
		}
	}
	return true
}

func (s *Sleigh) sort() {
	sort.Sort(s)

	//also need to sort each group
	for k, group := range s.groups {

		ints := []int{}
		for _, parcel := range group.parcels {
			ints = append(ints, int(parcel))
		}
		//sort them
		sort.Ints(ints)

		//write them back
		newParcels := []Parcel{}
		for _, val := range ints {

			newParcel := Parcel(val)
			newParcels = append(newParcels, newParcel)
		}
		s.groups[k].parcels = newParcels
	}
}

func (s *Sleigh) GetInfo() (label string, sizeFirstGroup int, quantumEntanglement int) {

	s.sort()

	sizeFirstGroup = len(s.groups[0].parcels)

	quantumEntanglement = 1
	for _, val := range s.groups[0].parcels {
		quantumEntanglement *= int(val)
	}

	chunks := []string{}
	for _, group := range s.groups {

		strParcels := []string{}
		for _, parcel := range group.parcels {

			strParcels = append(strParcels, parcel.toString())
		}

		chunk := strings.Join(strParcels, ",")
		chunks = append(chunks, chunk)

	}

	label = strings.Join(chunks, "|")

	return label, sizeFirstGroup, quantumEntanglement
}

// Need to be ab
func (s Sleigh) Len() int {
	return len(s.groups)
}
func (s Sleigh) Swap(i, j int) {
	s.groups[i], s.groups[j] = s.groups[j], s.groups[i]
}
func (s Sleigh) Less(i, j int) bool {

	if s.groups[i].len() < s.groups[j].len() {
		return true
	}

	if s.groups[i].len() > s.groups[j].len() {
		return false
	}

	//Need some quantum entanglement
	quanti := 1
	for _, parcel := range s.groups[i].parcels {
		quanti *= int(parcel)
	}
	quantj := 1
	for _, parcel := range s.groups[j].parcels {
		quantj *= int(parcel)
	}
	return quanti < quantj
}

func parse(input string) []int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	items := []int{}

	for scanner.Scan() {
		text := scanner.Text()

		int, _ := strconv.Atoi(text)
		items = append(items, int)
	}

	return items
}

type Result struct {
	sizeOfFirstPile     int
	quantumEntanglement int
}

func Run(input string, groupCount int) (quantumEntanglement int) {

	intParcels := parse(input)

	results := make(map[string]Result)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000; i++ {

		shuffle(intParcels)

		groups := []Group{}
		for i := 0; i < groupCount; i++ {
			groups = append(groups, Group{})
		}

		sleigh := Sleigh{groups}

		parcels := []Parcel{}
		for _, item := range intParcels {
			parcel := Parcel(item)
			parcels = append(parcels, parcel)
		}

		ok := sleigh.addParcels(parcels)
		if !ok {
			continue
		}

		label, sizeOne, qe := sleigh.GetInfo()

		result := Result{sizeOne, qe}
		results[label] = result
	}

	//The winner is the SMALLEST... or - the case of a tie - the one with the smallest qe

	sizeGroupOneMinimum := 10000000000000
	for _, result := range results {
		if result.sizeOfFirstPile < sizeGroupOneMinimum {
			sizeGroupOneMinimum = result.sizeOfFirstPile
		}
	}

	winningResults := []Result{}
	for _, result := range results {
		if result.sizeOfFirstPile == sizeGroupOneMinimum {
			winningResults = append(winningResults, result)
		}
	}

	QEmin := 10000000000000000
	for _, result := range winningResults {
		if result.quantumEntanglement < QEmin {
			QEmin = result.quantumEntanglement
		}
	}

	return QEmin
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
