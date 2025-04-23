package wizardry_test

import (
	"testing"

	"github.com/brainboxweb/advent/day22/wizardry"
	"github.com/stretchr/testify/require"
)

func TestBattles(t *testing.T) {

	isHard := false
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000, isHard) // remove bool

	require.Equal(t, 1269, minimumCost)
}

func TestBattlesHard(t *testing.T) {
	isHard := true
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000, isHard)

	require.Equal(t, 1309, minimumCost)
}
