package day10

import (
	"strconv"
)


func Run(input int, loops int) int {
	inputSlice := []int{}

	intString := strconv.Itoa(input)

	for _, integer := range intString {
		intint, _ := strconv.Atoi(string(integer))
		inputSlice = append(inputSlice, intint)
	}

	var result []int
	for i := 0; i < loops; i++ {

		if i == 0 {
			result = Say(inputSlice)
			continue
		}
		result = Say(result)
	}

	return len(result)
}

func Say(input []int) []int {
	output := []int{}
	reg := Register{}
	var regValue, regCount int

	for _, value := range input {

		regValue, regCount = reg.AddItem(value)
		if regValue > 0 {
			output = append(output, regCount)
			output = append(output, regValue)
		}
	}

	//Empty the register
	regValue, regCount = reg.Empty()
	if regValue > 0 {
		output = append(output, regCount)
		output = append(output, regValue)
	}

	return output
}

type Register struct {
	Value int
	Count int
}

//Add an integer.
//Returns either notthing... or the previous strng and count
func (r *Register) AddItem(item int) (int, int) {

	//On a run
	if item == r.Value {
		r.Count++
		return 0, 0
	}
	//Run has ended
	returnValue := r.Value
	returnCount := r.Count

	r.Value = item
	r.Count = 1

	return returnValue, returnCount
}

func (r *Register) Empty() (int, int) {

	returnValue := r.Value
	returnCount := r.Count

	r.Value = 0
	r.Count = 0

	return returnValue, returnCount
}
