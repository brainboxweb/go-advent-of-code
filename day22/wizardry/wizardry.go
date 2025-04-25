package wizardry

//	"fmt"

// Your spells are Magic Missile, Drain, Shield, Poison, and Recharge.
//
// Magic Missile costs 53 mana. It instantly does 4 damage.
// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.

type Effect struct {
	turns  int
	armor  int
	damage int
	cost   int
}

type Spell struct {
	name   string
	cost   int
	damage int
	heals  int
	effect Effect
}

/*
Weapons:    Cost  Damage  Armor

Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0
Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3
)
*/

type Player struct {
	name      string
	money     int
	hitPoints int
	damage    int
	spells    []Spell
}

type Battle struct {
	activeSpells map[Spell]int
}

// For the recurive function, I need to ditiniguish between losing and running out of spells
// Trying to cast two spells at the same time is considered a LOSE
func (b Battle) battle(player, boss Player) (isOk bool, winner Player) {
	turnIndex := 0
	for {
		turnIndex++

		ok, winner := b.takeTurn(turnIndex, &player, &boss)

		if !ok {
			return false, Player{}
		}

		if winner.name == "boss" || winner.name == "player" {
			return true, winner
		}
	}
}

func (b Battle) battleHard(player, boss Player) (isOk bool, winner Player) {
	turnIndex := 0
	for {
		turnIndex++

		player.hitPoints-- // Lose a hit point

		ok, winner := b.takeTurn(turnIndex, &player, &boss)

		if !ok {
			return false, Player{}
		}

		if winner.name == "boss" || winner.name == "player" {
			return true, winner
		}
	}
}

func (b *Battle) takeTurn(turnIndex int, player, boss *Player) (ok bool, winner Player) {
	ok = true

	if player.hitPoints <= 0 {
		return ok, *boss
	}

	// test for enough spells spell
	if len(player.spells) < turnIndex {
		return false, Player{} // Just false... because we want recursion to continue
	}

	result := b.castSpell(turnIndex, player, boss)
	if !result {
		// You lose!
		return ok, *boss
	}

	if player.money < 0 {
		return ok, *boss
	}

	if boss.hitPoints <= 0 {
		return ok, *player
	}

	// boss goes second
	b.attack(boss, player)
	if boss.hitPoints <= 0 {
		return ok, *player
	}
	if player.hitPoints <= 0 {
		return ok, *boss
	}

	return ok, Player{}
}

func (b *Battle) registerSpellEffect(spell Spell) bool {
	// nill map
	if b.activeSpells == nil {
		b.activeSpells = make(map[Spell]int)
	}
	if _, ok := b.activeSpells[spell]; ok {
		return false
	}
	b.activeSpells[spell] = spell.effect.turns
	return true
}

func (b *Battle) getSpells() []Spell {
	returnSpells := []Spell{}
	for spell := range b.activeSpells {
		returnSpells = append(returnSpells, spell)
		b.activeSpells[spell]--

		if b.activeSpells[spell] <= 0 {
			delete(b.activeSpells, spell)
		}
	}
	return returnSpells
}

func (b *Battle) castSpell(turnIndex int, attacker *Player, defender *Player) bool {
	damage := 0

	// Take the money
	// New spell: pay for it. register it.
	newSpell := attacker.spells[turnIndex-1]
	attacker.money -= attacker.spells[turnIndex-1].cost

	if attacker.money < 0 {
		return false
	}

	// Any spells to apply?
	activeSpells := b.getSpells()

	for _, spell := range activeSpells {
		attacker.money += spell.effect.cost
		damage += spell.effect.damage
	}

	if newSpell.effect != (Effect{}) {
		result := b.registerSpellEffect(newSpell)
		if !result {
			return false
		}
	} else {
		damage += newSpell.damage
		attacker.hitPoints += newSpell.heals
	}

	defender.hitPoints -= damage
	return true
}

// Damage dealt by an attacker each turn is equal to the attacker's damage score minus the defender's armor score.
// The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
func (b *Battle) attack(attacker, defender *Player) {
	defenderArmour := 0

	// effects apply on both turns
	activeSpells := b.getSpells()
	for _, spell := range activeSpells {
		defender.money += spell.effect.cost // NB - defender gets benefit
		attacker.hitPoints -= spell.effect.damage

		defenderArmour += spell.effect.armor

		if attacker.hitPoints <= 0 {
			return
		}
	}

	damage := attacker.damage - defenderArmour

	if damage < 1 {
		damage = 1
	}
	defender.hitPoints -= damage
}

func Battling(playerSpells []Spell, minimumCost int) int {
	allSpells := setSpells()

	if len(playerSpells) > 10 {
		return minimumCost
	}

	for _, spell := range allSpells {
		playerSpells := append(playerSpells, spell)

		player := Player{"player", 500, 50, 0, []Spell{}} // Hardcoded
		boss := Player{"boss", 0, 58, 9, []Spell{}}       // Puzzle input

		// Hit Points: 58
		// Damage: 9

		player.spells = playerSpells
		battle := Battle{}

		ok, winner := battle.battle(player, boss)

		if !ok {
			// Not finished... go deeper!
			deepMinimumCost := Battling(playerSpells, minimumCost)
			if deepMinimumCost < minimumCost {
				minimumCost = deepMinimumCost
			}
		}

		switch winner.name {
		case "player":
			// count the cost of spells
			totalCost := 0
			for _, spell := range playerSpells {
				totalCost += spell.cost
			}
			if totalCost < minimumCost {
				minimumCost = totalCost
			}
		case "boss":
			// Do nothing
		}
	}
	return minimumCost
}

func BattlingHard(playerSpells []Spell, minimumCost int) int {
	allSpells := setSpells()

	if len(playerSpells) > 10 {
		return minimumCost
	}

	for _, spell := range allSpells {
		playerSpells := append(playerSpells, spell)

		player := Player{"player", 500, 50, 0, []Spell{}} // Hardcoded
		boss := Player{"boss", 0, 58, 9, []Spell{}}       // Puzzle input

		// Hit Points: 58
		// Damage: 9

		player.spells = playerSpells
		battle := Battle{}

		ok, winner := battle.battleHard(player, boss)

		if !ok {
			// Not finished... go deeper!
			deepMinimumCost := BattlingHard(playerSpells, minimumCost)
			if deepMinimumCost < minimumCost {
				minimumCost = deepMinimumCost
			}
		}

		switch winner.name {
		case "player":
			// count the cost of spells
			totalCost := 0
			for _, spell := range playerSpells {
				totalCost += spell.cost
			}
			if totalCost < minimumCost {
				minimumCost = totalCost
			}
		case "boss":
			// Do nothing
		}
	}
	return minimumCost
}

func setSpells() []Spell {
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

	return spells
}
