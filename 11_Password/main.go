package main

import ()

func getNextPassword(oldPassword string) string {
	var newPassword string
	oldPasswordNumber := base26ToDecimal(oldPassword)
	for {
		oldPasswordNumber = oldPasswordNumber + 1
		newPassword = base26(oldPasswordNumber)
		if valid(newPassword) {
			return newPassword
		}
	}
}

func valid(input string) bool {
	return allowed(input) && doubleCount(input, 0) > 1 && rising(input)
}

func allowed(input string) bool {
	blacklist := []string{"i", "o", "l"}
	for _, char := range input {
		for _, black := range blacklist {
			if black == string(char) {

				return false
			}
		}
	}
	return true
}

func doubleCount(input string, total int) int {
	if len(input) < 2 {
		return total
	}
	output := ""
	previous := ""
	for k, character := range input {
		if k == 0 {
			//do nothing
		} else {
			//Same chars?
			if previous == string(character) {
				//its a double!
				total++
				output = input[k+1:]
				break
			}
		}
		previous = string(character)
	}
	return doubleCount(output, total)
}

func rising(input string) bool {
	run := 0
	previous := 0
	for _, char := range input {
		if int(char) == previous+1 {
			run++
			if run > 1 {
				return true
			}
		} else {
			run = 0 //reset
		}
		previous = int(char)
	}
	return false
}

func base26(num int) string {

	if num < 1 {
		panic("Not supported")
	}
	theString := ""
	for pos := 8; pos >= 0 && num > 0; pos-- {
		theString = string(int(rune(97))+(num%26)) + theString
		num = num / 26
	}
	return theString
}

func base26ToDecimal(input string) int {
	var output int
	position := len(input) - 1
	for _, char := range input {
		value := int(char) - 97
		positionalValue := value * simplePow(26, position)
		output += positionalValue
		position--
	}
	return output
}

func simplePow(x, y int) int {
	output := 1
	for i := 0; i < y; i++ {
		output = output * x
	}
	return output
}
