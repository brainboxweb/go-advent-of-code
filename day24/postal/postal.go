package postal

import (
	"bufio"
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type Parcel int

func (p Parcel) toString() string {
	return strconv.Itoa(int(p))
}

type Group struct {
	Parcels   []Parcel
	MaxWeight int
}

func (g Group) weight() (weight int) {
	for _, parcel := range g.Parcels {
		weight += int(parcel)
	}
	return weight
}

func (g Group) len() int {
	return len(g.Parcels)
}

func (g *Group) Add(p Parcel) bool {
	// Add... unless it exceeds maxweight
	if g.weight()+int(p) > g.MaxWeight {
		return false
	}

	g.Parcels = append(g.Parcels, p)
	return true
}

type Sleigh struct {
	Groups []Group
}

func (s *Sleigh) AddParcels(parcels []Parcel) bool {
	// Need descending order???
	// total weight:
	totalWeight := 0
	for _, parcel := range parcels {
		totalWeight += int(parcel)
	}
	groupWeight := totalWeight / len(s.Groups)
	// Inform the Groups
	for k := range s.Groups {
		s.Groups[k].MaxWeight = groupWeight
	}

	for _, parcel := range parcels {
		for k := range s.Groups {
			ok := s.Groups[k].Add(parcel)
			if ok {
				break // parcel placed successfully
			}
		}
	}

	// Check that ALL parcels have been placed.
	parcelCount := 0
	for _, group := range s.Groups {
		parcelCount += len(group.Parcels)
	}
	if parcelCount != len(parcels) {
		return false
	}

	// Check for balance
	for _, group := range s.Groups {
		if groupWeight != group.weight() {
			return false
		}
	}
	return true
}

func (s *Sleigh) sort() {
	sort.Sort(s)

	// also need to sort each group
	for k, group := range s.Groups {
		ints := []int{}
		for _, parcel := range group.Parcels {
			ints = append(ints, int(parcel))
		}
		// sort them
		sort.Ints(ints)

		// write them back
		newParcels := []Parcel{}
		for _, val := range ints {
			newParcel := Parcel(val)
			newParcels = append(newParcels, newParcel)
		}
		s.Groups[k].Parcels = newParcels
	}
}

func (s *Sleigh) GetInfo() (label string, sizeFirstGroup int, quantumEntanglement int) {
	s.sort()

	sizeFirstGroup = len(s.Groups[0].Parcels)

	quantumEntanglement = 1
	for _, val := range s.Groups[0].Parcels {
		quantumEntanglement *= int(val)
	}

	chunks := []string{}
	for _, group := range s.Groups {
		strParcels := []string{}
		for _, parcel := range group.Parcels {
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
	return len(s.Groups)
}
func (s Sleigh) Swap(i, j int) {
	s.Groups[i], s.Groups[j] = s.Groups[j], s.Groups[i]
}
func (s Sleigh) Less(i, j int) bool {
	if s.Groups[i].len() < s.Groups[j].len() {
		return true
	}

	if s.Groups[i].len() > s.Groups[j].len() {
		return false
	}

	// Need some quantum entanglement
	quanti := 1
	for _, parcel := range s.Groups[i].Parcels {
		quanti *= int(parcel)
	}
	quantj := 1
	for _, parcel := range s.Groups[j].Parcels {
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

		number, _ := strconv.Atoi(text)
		items = append(items, number)
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

		ok := sleigh.AddParcels(parcels)
		if !ok {
			continue
		}

		label, sizeOne, qe := sleigh.GetInfo()

		result := Result{sizeOne, qe}

		results[label] = result // try locking
	}

	// The winner is the SMALLEST... or - the case of a tie - the one with the smallest qe

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

	theQEmin := 10000000000000000
	for _, result := range winningResults {
		if result.quantumEntanglement < theQEmin {
			theQEmin = result.quantumEntanglement
		}
	}

	return theQEmin
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1) // #nosec G404
		a[i], a[j] = a[j], a[i]
	}
}
