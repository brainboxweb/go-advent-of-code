package postal

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
<<<<<<<< HEAD:_24_Packing/main.go
	parcels   []Parcel
	maxWeight int
========
	Parcels   []Parcel
	MaxWeight int
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
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
	//Add... unless it exceeds maxweight
<<<<<<<< HEAD:_24_Packing/main.go
	if g.weight()+int(p) > g.maxWeight {
========
	if g.weight()+int(p) > g.MaxWeight {
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
		return false
	}

	g.Parcels = append(g.Parcels, p)
	return true
}

type Sleigh struct {
	Groups []Group
}

<<<<<<<< HEAD:_24_Packing/main.go
func (s *Sleigh) addParcels(parcels []Parcel) bool {
========
func (s *Sleigh) AddParcels(parcels []Parcel) bool {
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
	//Need descending order???
	//tototal weufght:
	totalWeight := 0
	for _, parcel := range parcels {
		totalWeight += int(parcel)
	}
<<<<<<<< HEAD:_24_Packing/main.go
	groupWeight := totalWeight / len(s.groups)
	//Inform the Groups
	for k, _ := range s.groups {
		s.groups[k].maxWeight = groupWeight
	}

	for _, parcel := range parcels {
		for k, _ := range s.groups {
			ok := s.groups[k].Add(parcel)
========
	groupWeight := totalWeight / len(s.Groups)
	//Inform the Groups
	for k, _ := range s.Groups {
		s.Groups[k].MaxWeight = groupWeight
	}

	for _, parcel := range parcels {
		for k, _ := range s.Groups {
			ok := s.Groups[k].Add(parcel)
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
			if ok {
				break //parcel placed successfully
			}
		}
	}

	//Check that ALL parcels have been placed.
	parcelCount := 0
<<<<<<<< HEAD:_24_Packing/main.go
	for _, group := range s.groups {
		parcelCount += len(group.parcels)
========
	for _, group := range s.Groups {
		parcelCount += len(group.Parcels)
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
	}
	if parcelCount != len(parcels) {
		return false
	}

	//Check for balance
<<<<<<<< HEAD:_24_Packing/main.go
	for _, group := range s.groups {
========
	for _, group := range s.Groups {
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
		if groupWeight != group.weight() {
			return false
		}
	}
	return true
}

func (s *Sleigh) sort() {
	sort.Sort(s)

	//also need to sort each group
	for k, group := range s.Groups {

		ints := []int{}
<<<<<<<< HEAD:_24_Packing/main.go
		for _, parcel := range group.parcels {
========
		for _, parcel := range group.Parcels {
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
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

	//Need some quantum entanglement
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

<<<<<<<< HEAD:_24_Packing/main.go
	rand.Seed(time.Now().UnixNano())

========
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
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

<<<<<<<< HEAD:_24_Packing/main.go
		ok := sleigh.addParcels(parcels)
========
		ok := sleigh.AddParcels(parcels)
>>>>>>>> 91e6446 (Day 24):day24/postal/postal.go
		if !ok {
			continue
		}

		label, sizeOne, qe := sleigh.GetInfo()

		result := Result{sizeOne, qe}

		results[label] = result // try locking

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
