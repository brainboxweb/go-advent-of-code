package wizardry_test

import (
	"testing"

	"github.com/brainboxweb/advent/day22/wizardry"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	minimumCost := wizardry.Battling([]wizardry.Spell{}, 1000000)
	require.Equal(t, 1269, minimumCost)
}

func TestPart2(t *testing.T) {
	minimumCost := wizardry.BattlingHard([]wizardry.Spell{}, 1000000)
	require.Equal(t, 1309, minimumCost)
}
