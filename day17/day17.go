package day17

import (
	combos "github.com/mxschmitt/golang-combinations"
)

func New(capacities []int) Bottles {
	return Bottles{capacities: capacities}
}

type Bottles struct {
	capacities []int
}

func (bb Bottles) GetCombos(volume int) int {
	winnerCount := 0
	cc := combos.All(bb.capacities)
	for _, combo := range cc {
		if total(combo) == volume {
			winnerCount++
		}
	}
	return winnerCount
}

func (bb Bottles) GetLowestCombo(volume int) int {
	for i := 1; i < len(bb.capacities)+1; i++ {
		cc := combos.Combinations(bb.capacities, i)
		for _, combo := range cc {
			if total(combo) == volume {
				return i
			}
		}
	}
	return 0
}

func total(values []int) int {
	ret := 0
	for _, val := range values {
		ret += val
	}

	return ret
}
