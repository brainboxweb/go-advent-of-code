package wizardry_test

import (
	"testing"

	"github.com/brainboxweb/advent/day22/wizardry"
	"github.com/stretchr/testify/require"
)

func TestBattles(t *testing.T) {
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000)

	require.Equal(t, 1269, minimumCost)
}

func TestBattlesHard(t *testing.T) {
	minimumCost := wizardry.BattlingHard([]wizardry.Spell{}, 1000000)

	require.Equal(t, 1309, minimumCost)
}
