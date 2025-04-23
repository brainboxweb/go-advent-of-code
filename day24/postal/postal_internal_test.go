package postal

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestSortGroup(t *testing.T) {

	g1 := Group{[]Parcel{4, 3, 1, 7, 5}, 20}
	g2 := Group{[]Parcel{11, 9}, 20}
	g3 := Group{[]Parcel{10, 8, 2}, 20}
	gg := []Group{g1, g2, g3}
	sleigh := Sleigh{gg}

	//Expected
	gx1 := Group{[]Parcel{9, 11}, 20}
	gx2 := Group{[]Parcel{2, 8, 10}, 20}
	gx3 := Group{[]Parcel{1, 3, 4, 5, 7}, 20}
	ggx := []Group{gx1, gx2, gx3}
	expected := Sleigh{ggx}

	sleigh.sort()

	require.Equal(t, expected, sleigh)
}

func TestSortGroupWithQuantum(t *testing.T) {

	g1 := Group{[]Parcel{9, 5, 4, 2}, 20}
	g2 := Group{[]Parcel{11, 8, 1}, 20}
	g3 := Group{[]Parcel{10, 7, 3}, 20}
	gg := []Group{g1, g2, g3}
	sleigh := Sleigh{gg}

	//Expected
	gx1 := Group{[]Parcel{1, 8, 11}, 20} //88
	gx2 := Group{[]Parcel{3, 7, 10}, 20} //210
	gx3 := Group{[]Parcel{2, 4, 5, 9}, 20}
	ggx := []Group{gx1, gx2, gx3}
	expected := Sleigh{ggx}

	sleigh.sort()

	require.Equal(t, expected, sleigh)
}


func TestParse(t *testing.T) {
	input := `1
2
3
4
5
7
8
9
10
11`

	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	result := parse(input)
	require.Equal(t, expected, result)
}
