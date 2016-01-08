package main

import (
	"bufio"
	"bytes"
//	"math/rand"
	"sort"
	"strconv"
	"strings"
//	"time"
	"github.com/nightlyone/permutation"
	"reflect"
	"fmt"
)


/*
My modified algorithm is as follows:
Iteratively generate all combinations of length l for l from 1 to len(items)
For each such group of combinations, filter out groups for which their sum is not equal to sum(items) / n
Sort the remaining groups in increasing order of quantum entanglement
For each remaining group, if valid subdivision of items - group into n - 1 equal parts is possible, return that group immediately!
Otherwise, continue on to the next group of combinations
*/

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

func Run(input string, groupCount int) (quantumEntanglement int) {

//func Run(groupCount int, parcels []int) (qe int) {

	parcels := parse(input) //slice of ints
	maxPileSize := len(parcels)/groupCount

	totalWeight := 0
	for _, weight := range parcels{
		totalWeight += weight
	}


	//	maxPileSize :=


	fmt.Println(parcels)
	//NB maths  - think it will be okay

//	combos := [][]int{}

	fmt.Println("COMBOS")
//	//start with all combinations
//	for _, val := range parcels{
//		startSlice := []int{val}
//		subCombo := combine(startSlice, parcels, maxPileSize)
//		combos = append(combos, subCombo...)
//	}



	combos := combine([]int{}, parcels, maxPileSize)



		fmt.Println(combos)


	fmt.Println("CANDIDATES")
	//Check the totals
	candidates := [][]int{}
	for _, set := range combos {
		total := 0
		for _, val := range set{
			total += val
		}
		if total == totalWeight/groupCount {
			candidates = append(candidates, set)
		}
	}

	//	fmt.Println(candidates)
	//	//Get the



	fmt.Println("POSSIBLES")
	possibles := Piles{}
	for _, set := range candidates {
		qe := 1
		for _, val := range set {
			qe *= val
		}
		pile := Pile{set, len(set), qe}

		possibles.piles = append(possibles.piles, pile)
	}


	sort.Sort(possibles)

	fmt.Println(len(possibles.piles))


	i := maxPileSize/(groupCount)  // too big?

	fmt.Println(i)

	for {


		for _, pile := range possibles.piles {
			if pile.count != i {
				continue
			}
			//return IMMEDIATELY
			return pile.qe
		}

		i ++
	}

	return quantumEntanglement

}


func combine(currentList, remainingList []int, maxLength int) (combinations [][]int) {

	for k, val := range remainingList{

		currentListCopy := make([]int, len(currentList))
		copy(currentListCopy, currentList)

		currentListCopy = append(currentListCopy, val)
		combinations = append(combinations, currentList)


		if len(currentListCopy) <= maxLength {
			//Go deeper
			newCombos := combine(currentListCopy, remainingList[k+1:], maxLength)
			combinations = append(combinations, newCombos...)
		}

	}
	return combinations
}


type Pile struct {
	parcels   []int
	count int
	qe int
}

type Piles struct {
	piles []Pile
}

func (p Piles) Len() int {
	return len(p.piles)
}
func (p Piles) Swap(i, j int) {
	p.piles[i], p.piles[j] = p.piles[j], p.piles[i]
}
func (p Piles) Less(i, j int) bool {
	return p.piles[i].qe < p.piles[j].qe
}




func recurse(targetWeight, maxCount int, bigList, currentList []int, winners [][]int ) [][]int {

//	fmt.Println("\n\nRECURSE>>>>>>>>>>>")
	for _,val := range bigList {

		currentList = append(currentList, val)

		fmt.Println(currentList)

		//testWeight
		totalWeight := 0
		for _, parcelWeight := range currentList {
			totalWeight += parcelWeight
			switch {
			case totalWeight == targetWeight:
				fmt.Println("WINNER!!!!")
			case totalWeight > targetWeight:
				fmt.Println("BUST!!!!")
				//FAILED!!!
				return winners
			default:
				//Go again

				if len(currentList) < maxCount {
					fmt.Println("RECURSE!!!!")
					recurse(targetWeight, maxCount, bigList, currentList, winners)
				}
			}
		}
	}
	return winners

}





//
func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func main() {
	orig := []int{11, 10, 9, 8, 7, 5, 4, 3, 2, 1}
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		fmt.Println(getPerm(orig, p))
	}
}












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

	//	Need smallest combination that adds up to (sum/piles) = 20
	groupTotal := 0
	for _, parcel := range parcels {

		groupTotal += int(parcel)

		switch {
		case groupTotal == groupWeight:
			return true
		case groupTotal > groupWeight:
			return false
		}


//		for k, _ := range s.groups {
//			ok := s.groups[k].Add(parcel)
//			if ok {
//				break //parcel placed successfully
//			}
//		}
	}
//
//	//Check that ALL parcels have been placed.
//	parcelCount := 0
//	for _, group := range s.groups {
//		parcelCount += len(group.parcels)
//	}
//	if parcelCount != len(parcels) {
//		return false
//	}
//
//	//Check for balance
//	for _, group := range s.groups {
//		if groupWeight != group.weight() {
//			return false
//		}
//	}
	return false
}



//
//11,10,9,8,7,5,4,3,2,1  - sum is 60
//
//Need smallest combination that adds up to (sum/3) = 20
//
//try 11 - too small
//try 11, 10 - too small
//try 11, 10, 9 - winner!

/*
Python. Just need to find the smallest combination of numbers that adds up to the sum of the weights divided by 3
 (or 4 for part 2) since you need 3 (or 4) equal groups. Of the combinations that satisfy that condition, find the
 minimum quantum entanglement.
day = 24

from functools import reduce
from itertools import combinations
from operator import mul

wts = [int(x) for x in get_input(day).split('\n')]

def day24(num_groups):
group_size = sum(wts) // num_groups
for i in range(len(wts)):
qes = [reduce(mul, c) for c in combinations(wts, i)
if sum(c) == group_size]
if qes:
return min(qes)

print(day24(3))
print(day24(4))
*/









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

//Need to be ab
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

func parseOld(input string) []int {
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

func RunOLD(input string, groupCount int) (quantumEntanglement int) {

	intParcels := parse(input)

//	sort into descending order
	sort.Sort(sort.Reverse(sort.IntSlice(intParcels))) //Pack biggest first


	results := make(map[string]Result)

	groups := []Group{}
	for i := 0; i < groupCount; i++ {
		groups = append(groups, Group{})
	}

	fmt.Println(groups)

	sleigh := Sleigh{groups}



	keys := intSlice{}
	for _, val := range intParcels {
		keys = append(keys, val)
	}
	perm := permutation.New(keys)
	// iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {

//		fmt.Println(seq)
//
//		fmt.Printf(".")
		indexes := seqToSlice(seq)

//		fmt.Println(results)

		//This is madness!!
		parcels := []Parcel{}



		for _, val := range indexes {
			parcel := Parcel(val)
			parcels = append(parcels, parcel)
		}



		ok := sleigh.addParcels(parcels)
		if !ok {
			continue
		}

		fmt.Println(parcels)


		label, sizeOne, qe := sleigh.GetInfo()

		result := Result{sizeOne, qe}
		results[label] = result

		fmt.Println(results)

	}

	fmt.Println(results)

//	//try removing a different package on each loop
//	for i := 0 ; i < len(intParcels); i++ {
////
////		intParcelsCopy := make([]int, len(intParcels))
////
////		if i == 0 {
////			intParcelsCopy = intParcels
////		} else {
////			//switch the parcels around
////			copy(intParcelsCopy, intParcels)
////			intParcelsCopy = append(intParcelsCopy[i:], intParcelsCopy[:i]...) //worked for most cases
////		}
//
//		// Permutation of simple slice
//
//
//
//
//
//		parcels := []Parcel{}
//		for _, item := range intParcelsCopy {
//			parcel := Parcel(item)
//			parcels = append(parcels, parcel)
//		}
//
//		ok := sleigh.addParcels(parcels)
//		if !ok {
////			continue
//		}
//
//		label, sizeOne, qe := sleigh.GetInfo()
//
//		result := Result{sizeOne, qe}
//		results[label] = result
//
//	}

	fmt.Println(results)

// WORKED... BUT RELIES ON RAMDOMNESS
//	//Okay... this is silly
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < 100000; i++ {
//
//		shuffle(intParcels)
//
//		groups := []Group{}
//		for i := 0; i < groupCount; i++ {
//			groups = append(groups, Group{})
//		}
//
//		sleigh := Sleigh{groups}
//
//		parcels := []Parcel{}
//		for _, item := range intParcels {
//			parcel := Parcel(item)
//			parcels = append(parcels, parcel)
//		}
//
//		ok := sleigh.addParcels(parcels)
//		if !ok {
//			continue
//		}
//
//		label, sizeOne, qe := sleigh.GetInfo()
//
//		result := Result{sizeOne, qe}
//		results[label] = result
//	}

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
//
//func shuffle(a []int) {
//	for i := range a {
//		j := rand.Intn(i + 1)
//		a[i], a[j] = a[j], a[i]
//	}
//}


//
//func (a []int) Len() int      { return len(d) }
//func (d distances) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
//func (d distances) Less(i, j int) bool {
//	if d[i].Distance > d[j].Distance { //reverse search
//		return true
//	} else {
//		return false
//	}
//}


//67601337186  //Unblanced answers
//11846773891  //expected answer
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
