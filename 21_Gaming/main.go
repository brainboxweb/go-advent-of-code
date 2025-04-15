package main

import (
//	"fmt"
)

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
	hitPoints int
	damage    int
	armour    int
}

type Weapon struct {
	name   string
	cost   int
	damage int
}

type Armour struct {
	name   string
	cost   int
	armour int
}

type Ring struct {
	name   string
	cost   int
	damage int
	armour int
}

func play() (minimumCost, maximumCost int) {

	//	Hit Points: 109
	//	Damage: 8
	//	Armor: 2
	player := Player{"player", 100, 0, 0}
	boss := Player{"boss", 109, 8, 2}

	minimumCost = 1000000
	maximumCost = 0

	weapons := weapons()
	for _, weapon := range weapons {

		armour := armour()

		for _, armourItem := range armour {
			armourSlice := []Armour{armourItem}
			theRings := rings()
			for _, ring := range theRings {

				ringSlice := []Ring{ring}

				winner, cost := doIt(player, boss, weapon, armourSlice, ringSlice)
				if winner.name == "player" && cost < minimumCost {
					minimumCost = cost
				}
				if winner.name == "boss" && cost > maximumCost {
					maximumCost = cost
				}

				theRings2 := theRings
				//mix in another ring
				for _, ring2 := range theRings2 {
					if ring == ring2 {
						continue
					}
					ringSlice2 := []Ring{ring, ring2}
					winner, cost := doIt(player, boss, weapon, armourSlice, ringSlice2)
					if winner.name == "player" && cost < minimumCost {
						minimumCost = cost
					}
					if winner.name == "boss" && cost > maximumCost {
						maximumCost = cost
					}
				}
			}
		}
	}

	return minimumCost, maximumCost
}

func doIt(player, boss Player, weapon Weapon, armour []Armour, rings []Ring) (winner Player, cost int) {

	//	fmt.Println("\n\n\n--------------------\n",weapon, armour, rings )
	cost = 0

	//Weapon
	player.damage += weapon.damage
	cost += weapon.cost

	//Armour
	for _, armourItem := range armour {
		player.armour += armourItem.armour
		cost += armourItem.cost
	}

	//Rings
	for _, ring := range rings {
		player.armour += ring.armour
		player.damage += ring.damage
		cost += ring.cost
	}

	winner = battle(player, boss)

	return winner, cost
}

func battle(player, boss Player) (winner Player) {

	for {
		//Player goes first
		attack(&player, &boss)
		if boss.hitPoints <= 0 {
			return player
		}

		//boss goes second
		attack(&boss, &player)
		if player.hitPoints <= 0 {
			return boss
		}
	}
}

//	Damage dealt by an attacker each turn is equal to the attacker's damage score minus the defender's armor score.
//	The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
func attack(attacker, defender *Player) {
	damage := attacker.damage - defender.armour
	if damage < 1 {
		damage = 1
	}
	defender.hitPoints -= damage
	//	fmt.Printf("\nThe %s deals %d-%d = %d damage; the %s goes down to %d hit points.", attacker.name, attacker.damage, defender.armour, damage, defender.name, defender.hitPoints)
}

func weapons() []Weapon {

	dagger := Weapon{"Dagger", 8, 4}
	shortsword := Weapon{"Shortsword", 10, 5}
	warhammer := Weapon{"Warhammer", 25, 6}
	longsword := Weapon{"Longsword", 40, 7}
	greataxe := Weapon{"Greataxe", 74, 8}

	weapons := []Weapon{
		dagger,
		shortsword,
		warhammer,
		longsword,
		greataxe,
	}
	return weapons
}

func armour() []Armour {

	dummy := Armour{"DUMMY", 0, 0}
	leather := Armour{"Leather", 13, 1}
	chainmail := Armour{"Chainmail", 31, 2}
	splintmail := Armour{"Splintmail", 53, 3}
	bandedmail := Armour{"Bandedmail", 75, 4}
	platemail := Armour{"Platemail", 102, 5}

	armour := []Armour{
		dummy,
		leather,
		chainmail,
		splintmail,
		bandedmail,
		platemail,
	}
	return armour
}

func rings() []Ring {

	da0 := Ring{"DUMMY", 0, 0, 0}
	da1 := Ring{"Damage +1", 25, 1, 0}
	da2 := Ring{"Damage +2", 50, 2, 0}
	da3 := Ring{"Damage +3", 100, 3, 0}
	de1 := Ring{"Defense +1", 20, 0, 1}
	de2 := Ring{"Defense +2", 40, 0, 2}
	de3 := Ring{"Defense +3", 80, 0, 3}

	rings := []Ring{
		da0,
		da1,
		da2,
		da3,
		de1,
		de2,
		de3,
	}
	return rings
}
