package day21

import "github.com/brainboxweb/advent/day21/battleground"

func PlayToWin() int { //  <-- Part 1
	minimumCost, _ := battleground.Play()
	return minimumCost
}

func PlayToLose() int { //  <-- Part 2
	_, maximumCost := battleground.Play()
	return maximumCost
}
