package wizardry_test

import (
	"testing"

	"github.com/brainboxweb/advent/day22/wizardry"
	"github.com/stretchr/testify/require"
)

func Part1(t *testing.T) {
	isHard := false
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000, isHard)
	require.Equal(t, 1269, minimumCost)
}

func Part2(t *testing.T) {
	isHard := true
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000, isHard)
	require.Equal(t, 1309, minimumCost)
}
