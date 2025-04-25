package day20

func GetHouseNumber(presentCount int, presentsPerElf int, limitPerElf int) int {
	house := make(map[int]int)
	loopLimit := presentCount / presentsPerElf

	for i := 1; i < loopLimit; i++ { // loop through the elves
		deliveryCount := 0                  // Count the delivery count for each elf
		for j := i; j < loopLimit; j += i { // loop through the houses in steps of the current "elf"
			deliveryCount++
			if limitPerElf > 0 && deliveryCount > limitPerElf { // For case where count is limited
				break
			}
			house[j] += i * presentsPerElf
		}
	}

	// Find the lowest house number
	lowestHouseNumber := 100000000000000
	for houseNumber, val := range house {
		if val >= presentCount {
			if houseNumber < lowestHouseNumber {
				lowestHouseNumber = houseNumber
			}
		}
	}
	return lowestHouseNumber
}
