package postal_test

import (
	"testing"

	"github.com/brainboxweb/advent/day24/postal"
	"github.com/stretchr/testify/require"
)

func TestAddParcels(t *testing.T) {
	parcels := []postal.Parcel{11, 10, 9, 8, 7, 5, 4, 3, 2, 1}

	group1 := postal.Group{}
	group2 := postal.Group{}
	group3 := postal.Group{}
	groups := []postal.Group{group1, group2, group3}

	sleigh := postal.Sleigh{groups}
	sleigh.AddParcels(parcels)

	// Expected
	g1 := postal.Group{[]postal.Parcel{11, 9}, 20}
	g2 := postal.Group{[]postal.Parcel{10, 8, 2}, 20}
	g3 := postal.Group{[]postal.Parcel{7, 5, 4, 3, 1}, 20}
	gg := []postal.Group{g1, g2, g3}
	expected := postal.Sleigh{gg}

	require.Equal(t, expected, sleigh)
}

func TestGetInfo(t *testing.T) {
	g1 := postal.Group{[]postal.Parcel{1, 8, 11}, 20}
	g2 := postal.Group{[]postal.Parcel{3, 7, 10}, 20}
	g3 := postal.Group{[]postal.Parcel{2, 4, 5, 9}, 20}
	gg := []postal.Group{g1, g2, g3}
	sleigh := postal.Sleigh{gg}

	expectedLabel := "1,8,11|3,7,10|2,4,5,9"
	expectedSizeFirstGroup := 3
	expectedQuantumEntanglement := 88

	label, sizeFirstGroup, quantumEntanglement := sleigh.GetInfo()

	require.Equal(t, expectedLabel, label)
	require.Equal(t, expectedSizeFirstGroup, sizeFirstGroup)
	require.Equal(t, expectedQuantumEntanglement, quantumEntanglement)
}

// // Damn. It's unstable

func TestRun(t *testing.T) {
	var tests = []struct {
		input      string
		groupCount int
		expected   int
	}{
		{
			`1
2
3
4
5
7
8
9
10
11`,
			3,
			99,
		},
	}
	for _, test := range tests {
		if actual := postal.Run(test.input, test.groupCount); actual != test.expected {
			t.Errorf("Convert(%q) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}
