package day1

import (
	"errors"
)

func FinalFloor(input string) int {

	floor := 0
	for _, r := range input {
		c := string(r)
		switch c {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
}

func Basement(input string) (int, error) {

	floor := 0
	count := 0
	for _, r := range input {
		count++
		c := string(r)
		switch c {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			return count, nil
		}
	}
	return 0, errors.New("Didn't make it to the basement")
}
