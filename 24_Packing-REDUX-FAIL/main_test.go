package main

import (
//	"github.com/stretchr/testify/require"
	"testing"
)


//func TestPlay(t *testing.T) {
//
//	groups := 3
//	parcels := []int{11, 10, 9, 8, 7, 5, 4, 3, 2, 1}
//
//
//	qe := play(groups, parcels)
//	expected := 98
//	require.Equal(t, expected, qe)
//}




//func TestAddParcels(t *testing.T) {
//
//	parcels := []Parcel{11, 10, 9, 8, 7, 5, 4, 3, 2, 1}
//
//	group1 := Group{}
//	group2 := Group{}
//	group3 := Group{}
//	groups := []Group{group1, group2, group3}
//	sleigh := Sleigh{groups}
//	sleigh.addParcels(parcels)
//
//	//Expected
//	g1 := Group{[]Parcel{11, 9}, 20}
//	g2 := Group{[]Parcel{10, 8, 2}, 20}
//	g3 := Group{[]Parcel{7, 5, 4, 3, 1}, 20}
//	gg := []Group{g1, g2, g3}
//	expected := Sleigh{gg}
//
//	require.Equal(t, expected, sleigh)
//}

//
//
//
//func TestSortGroup(t *testing.T) {
//
//	g1 := Group{[]Parcel{4, 3, 1, 7, 5}, 20}
//	g2 := Group{[]Parcel{11, 9}, 20}
//	g3 := Group{[]Parcel{10, 8, 2}, 20}
//	gg := []Group{g1, g2, g3}
//	sleigh := Sleigh{gg}
//
//	//Expected
//	gx1 := Group{[]Parcel{9, 11}, 20}
//	gx2 := Group{[]Parcel{2, 8, 10}, 20}
//	gx3 := Group{[]Parcel{1, 3, 4, 5, 7}, 20}
//	ggx := []Group{gx1, gx2, gx3}
//	expected := Sleigh{ggx}
//
//	sleigh.sort()
//
//	require.Equal(t, expected, sleigh)
//}
//
//func TestSortGroupWithQuantum(t *testing.T) {
//
//	g1 := Group{[]Parcel{9, 5, 4, 2}, 20}
//	g2 := Group{[]Parcel{11, 8, 1}, 20}
//	g3 := Group{[]Parcel{10, 7, 3}, 20}
//	gg := []Group{g1, g2, g3}
//	sleigh := Sleigh{gg}
//
//	//Expected
//	gx1 := Group{[]Parcel{1, 8, 11}, 20} //88
//	gx2 := Group{[]Parcel{3, 7, 10}, 20} //210
//	gx3 := Group{[]Parcel{2, 4, 5, 9}, 20}
//	ggx := []Group{gx1, gx2, gx3}
//	expected := Sleigh{ggx}
//
//	sleigh.sort()
//
//	require.Equal(t, expected, sleigh)
//}
//
//func TestGetInfo(t *testing.T) {
//
//	g1 := Group{[]Parcel{1, 8, 11}, 20}
//	g2 := Group{[]Parcel{3, 7, 10}, 20}
//	g3 := Group{[]Parcel{2, 4, 5, 9}, 20}
//	gg := []Group{g1, g2, g3}
//	sleigh := Sleigh{gg}
//
//	expectedLabel := "1,8,11|3,7,10|2,4,5,9"
//	expectedSizeFirstGroup := 3
//	expectedQuantumEntanglement := 88
//
//	label, sizeFirstGroup, quantumEntanglement := sleigh.GetInfo()
//
//	require.Equal(t, expectedLabel, label)
//	require.Equal(t, expectedSizeFirstGroup, sizeFirstGroup)
//	require.Equal(t, expectedQuantumEntanglement, quantumEntanglement)
//}
//
//func TestParse(t *testing.T) {
//	input := `1
//2
//3
//4
//5
//7
//8
//9
//10
//11`
//
//	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
//	result := parse(input)
//	require.Equal(t, expected, result)
//}
//
var tests = []struct {
	input      string
	groupCount int
	expected   int
}{
	{
		`11
10
9
8
7
5
4
3
2
1`,
		3,
		99,
	},
//	{
//		`1
//2
//3
//4
//5
//7
//8
//9
//10
//11`,
//		3,
//		99,
//	},
	{
		day24data,
		3,
		11846773891,
	},
//	{
//		day24data,
//		4,
//		80393059,
//	},
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		if actual := Run(test.input, test.groupCount); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day24data = `1
2
3
7
11
13
17
19
23
31
37
41
43
47
53
59
61
67
71
73
79
83
89
97
101
103
107
109
113`

//89,103,107,109 = 408
//23,41,43,47,53,59,73,79,101 = 519
//1,2,3,7,11,13,17,19,31,37,61,67,71,83,97
