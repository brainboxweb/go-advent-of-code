package wizardry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTakeTurn(t *testing.T) {
	battle := Battle{}

	player := Player{"player", 250, 10, 0, []Spell{}}
	boss := Player{"boss", 0, 13, 8, []Spell{}}

	poison := Spell{
		"poison",
		173,
		0,
		0,
		Effect{
			6,
			0,
			3,
			0,
		},
	}

	magicMissile := Spell{
		"magicMissile",
		53,
		4,
		0,
		Effect{},
	}

	spells := []Spell{
		poison,
		magicMissile,
		magicMissile,
	}

	player.spells = spells

	battle.takeTurn(1, &player, &boss)

	battle.takeTurn(2, &player, &boss)

	require.Equal(t, 0, boss.hitPoints)
}

func TestTakeTurnAgain(t *testing.T) {
	player := Player{"player", 250, 10, 0, []Spell{}}
	boss := Player{"boss", 0, 14, 8, []Spell{}}

	recharge := Spell{
		"recharge",
		229,
		0,
		0,
		Effect{
			5,
			0,
			0,
			101,
		},
	}

	sheild := Spell{
		"sheild",
		113,
		0,
		0,
		Effect{
			6,
			7,
			0,
			0,
		},
	}

	drain := Spell{
		"drain",
		73,
		2,
		2,
		Effect{},
	}

	poison := Spell{
		"poison",
		173,
		0,
		0,
		Effect{
			6,
			0,
			3,
			0,
		},
	}

	magicMissile := Spell{
		"magicMissile",
		53,
		4,
		0,
		Effect{},
	}

	spells := []Spell{
		recharge,
		sheild,
		drain,
		poison,
		magicMissile,
	}

	player.spells = spells

	battle := Battle{}

	_, winner := battle.battle(player, boss)

	require.Equal(t, "player", winner.name)
}
