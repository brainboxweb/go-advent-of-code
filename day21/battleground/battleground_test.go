package battleground

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//
// For example, suppose the player has 10 hit points and 250 mana, and that the boss has 13 hit points and 8 damage:
//
// -- Player turn --
// - Player has 10 hit points, 0 armor, 250 mana
// - Boss has 13 hit points
// Player casts Poison.
//
// -- Boss turn --
// - Player has 10 hit points, 0 armor, 77 mana
// - Boss has 13 hit points
// Poison deals 3 damage; its timer is now 5.
// Boss attacks for 8 damage.
//
// -- Player turn --
// - Player has 2 hit points, 0 armor, 77 mana
// - Boss has 10 hit points
// Poison deals 3 damage; its timer is now 4.
// Player casts Magic Missile, dealing 4 damage.
//
// -- Boss turn --
// - Player has 2 hit points, 0 armor, 24 mana
// - Boss has 3 hit points
// Poison deals 3 damage. This kills the boss, and the player wins.

func TestBattle(t *testing.T) {
	player1 := Player{"player", 8, 5, 5}
	player2 := Player{"boss", 12, 7, 2}
	winner := battle(player1, player2)
	expected := Player{"player", 2, 5, 5} // Player 1 at the end
	require.Equal(t, expected, winner)
}

// Damage dealt by an attacker each turn is equal to the attacker's damage score minus the defender's armor score. An attacker always does at least 1 damage.
//  So, if the attacker has a damage score of 8, and the defender has an armor score of 3, the defender loses 5 hit points.
//  If the defender had an armor score of 300, the defender would still lose 1 hit point.

func TestDamage(t *testing.T) {
	attacker := Player{"attack", 100, 8, 0}
	defender := Player{"defend", 100, 0, 3}
	attack(&attacker, &defender)
	expectedHitPoints := 95
	require.Equal(t, expectedHitPoints, defender.hitPoints)
}

func TestDamage2(t *testing.T) {
	attacker := Player{"attack", 100, 8, 0}
	defender := Player{"defend", 100, 0, 300}
	attack(&attacker, &defender)
	expectedHitPoints := 99
	require.Equal(t, expectedHitPoints, defender.hitPoints)
}

func TestPlayToWin(t *testing.T) { //  <-- Part 1
	cost, _ := Play()
	expected := 111
	require.Equal(t, expected, cost)
}

func TestPlayToLose(t *testing.T) { //  <-- Part 2
	_, cost := Play()
	expected := 188
	require.Equal(t, expected, cost)
}
