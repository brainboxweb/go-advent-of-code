package main

import(
//	"fmt"
)


//Your spells are Magic Missile, Drain, Shield, Poison, and Recharge.
//
//Magic Missile costs 53 mana. It instantly does 4 damage.
//Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
//Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
//Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
//Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.


type Effect struct{
	turns int
	armor int
	damage int
	cost int
}

type  Spell struct{
	name string
	cost int
	damage int
	heals int
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


type Player struct{
	name string
	money int
	hitPoints int
	damage int
	spells []Spell
}

type Weapon struct{
	name string
	cost int
	damage int
}

type Armour struct{
	name string
	cost int
	armour int
}

type Ring struct{
	name string
	cost int
	damage int
	armour int
}

type Battle struct{
	activeSpells map[Spell]int
}
//
//func weapons() []Weapon {
//
//	dagger := Weapon{"Dagger",          8,     4}
//	shortsword := Weapon{"Shortsword", 10,     5}
//	warhammer := Weapon{"Warhammer",   25,     6}
//	longsword := Weapon{"Longsword",   40,     7}
//	greataxe := Weapon{"Greataxe",     74,     8}
//
//	weapons := []Weapon{
//		dagger,
//		shortsword,
//		warhammer,
//		longsword,
//		greataxe,
//	}
//	return weapons
//}
//
//func armour() []Armour {
//
//	dummy := Armour{"DUMMY",            0,  0}
//	leather := Armour{"Leather",       13,  1}
//	chainmail := Armour{"Chainmail",   31,  2}
//	splintmail := Armour{"Splintmail", 53,  3}
//	bandedmail := Armour{"Bandedmail", 75,  4}
//	platemail := Armour{"Platemail",   102, 5}
//
//	armour := []Armour{
//		dummy,
//		leather,
//		chainmail,
//		splintmail,
//		bandedmail,
//		platemail,
//	}
//	return armour
//}
//
//func rings() []Ring {
//
//	da0 := Ring{"DUMMY",         0,     0,       0}
//	da1 := Ring{"Damage +1",    25,     1,       0}
//	da2 := Ring{"Damage +2",    50,     2,       0}
//	da3 := Ring{"Damage +3",   100,     3,       0}
//	de1 := Ring{"Defense +1",   20,     0,       1}
//	de2 := Ring{"Defense +2",   40,     0,       2}
//	de3 := Ring{"Defense +3",   80,     0,       3}
//
//	rings := []Ring{
//		da0,
//		da1,
//		da2,
//		da3,
//		de1,
//		de2,
//		de3,
//	}
//	return rings
//}

//For the recurive function, I need to ditiniguish between losing and running out of spells
//Trying to cast two spells at the same time is considered a LOSE
func (b Battle) battle(player, boss Player, isHard bool) (isOk bool, winner Player) {

	turnIndex := 0
	for {
		turnIndex++

		if isHard == true {
			player.hitPoints = player.hitPoints - 1  //Lose a hit point
		}

		ok, winner := b.takeTurn(turnIndex, &player, &boss)

		if ok == false {
			return false, Player{}
		}

		if winner.name == "boss" || winner.name == "player" {
			return true, winner
		}

	}
	return false, Player{}
}


func (b *Battle) takeTurn(turnIndex int, player, boss *Player) (ok bool, winner Player){

	ok = true

	if player.hitPoints <= 0 {
		return ok, *boss
	}

	//test for enough spells spell
	if len(player.spells) < turnIndex {
		return false, Player{}  //Just false... because we want recursion to continue
	}

	result := b.castSpell(turnIndex, player, boss)
	if result == false {
		//You lose!
		return ok, *boss
	}

	if player.money < 0 {
		return ok, *boss
	}

	if boss.hitPoints <= 0 {
		return ok, *player
	}

	//boss goes second
	b.attack(boss, player)
	if boss.hitPoints <= 0 {
		return ok, *player
	}
	if player.hitPoints <= 0 {
		return ok, *boss
	}

	return ok, Player{}
}


func (b *Battle) registerSpellEffect(spell Spell) bool{
	//nill map
	if b.activeSpells == nil {
		b.activeSpells = make(map[Spell]int)
	}
	if _, ok := b.activeSpells[spell]; ok {
//		panic("cannot have two at the same time")
		return false
	}
	b.activeSpells[spell] = spell.effect.turns
	return true
}

func (b *Battle) getSpells()  []Spell {

	returnSpells := []Spell{}
	for spell, _ := range b.activeSpells {
		returnSpells = append(returnSpells, spell)
		b.activeSpells[spell]--

		if b.activeSpells[spell] <= 0 {
			delete(b.activeSpells, spell)
		}
	}
	return returnSpells
}


func  (b *Battle) castSpell(turnIndex int, attacker *Player, defender *Player) bool {

	damage := 0

	//Take the money
	//New spell: pay for it. register it.
	newSpell := attacker.spells[turnIndex-1]
	attacker.money -= attacker.spells[turnIndex-1].cost


	if attacker.money < 0 {
		return false
	}

	//Any spells to apply?
	activeSpells := b.getSpells()

	for _, spell := range activeSpells {
		attacker.money += spell.effect.cost
		damage += spell.effect.damage
	}

	if newSpell.effect != (Effect{}) {
		result := b.registerSpellEffect(newSpell)
		if result == false {
			return false
		}
	} else {
		damage += newSpell.damage
		attacker.hitPoints += newSpell.heals
	}

	defender.hitPoints -=  damage
	return true
}


//	Damage dealt by an attacker each turn is equal to the attacker's damage score minus the defender's armor score.
//	The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
func   (b *Battle)  attack(attacker, defender *Player){

	defenderArmour := 0

	//effects apply on both turns
	activeSpells := b.getSpells()
	for _, spell := range activeSpells {

		defender.money += spell.effect.cost //NB - defender gets benefit
		attacker.hitPoints  -= spell.effect.damage

		defenderArmour += spell.effect.armor

		if attacker.hitPoints <= 0 {
			return
		}

	}

	damage := attacker.damage - defenderArmour

	if damage < 1 {
		damage = 1
	}
	defender.hitPoints -=  damage
}




func Battling (allSpells, playerSpells []Spell, minimumCost int, isHard bool) int {

	if len(playerSpells) > 10 {
		return minimumCost
	}

	for _, spell := range allSpells{

		newSpells := append(playerSpells, spell)

		player := Player{"player", 500, 50, 0, []Spell{}}
		boss := Player{"boss", 0, 55, 8, []Spell{}}
		player.spells = newSpells
		battle := Battle{}

		ok, winner := battle.battle(player, boss, isHard)
		if ok == false {
			//Not finished... go deeper!
			deepMinimumCost := Battling(allSpells, newSpells, minimumCost, isHard)
			if deepMinimumCost < minimumCost {
				minimumCost = deepMinimumCost
			}
		}

		switch winner.name{
		case "player":
			//count the cost of spells
			totalCost := 0
			for _, spell := range newSpells {
				totalCost += spell.cost
			}
			if totalCost < minimumCost {
				minimumCost = totalCost
			}
		case "boss":
			//Do nothing
		}

	}
	return minimumCost
}
