package day25

func run(start, targetX, targetY int) int {
	var row, col int
	if targetX > targetY {
		row = targetX * 2
		col = targetX * 2
	} else {
		row = targetY * 2
		col = targetY * 2
	}

	coords := make([][]int, row)
	e := make([]int, row*col)
	for i := range coords {
		coords[i] = e[i*col : (i+1)*col]
	}

	x := 1
	y := 1
	coords[x][y] = int(start)
	var lastValue int
	lastValue = int(start)
	for {
		// taking the previous one, multiplying it by 252533, and then keeping the remainder from dividing that value by 33554393.
		newValue := (lastValue * 252533) % 33554393
		x, y = step(x, y)

		if x == targetX && y == targetY {
			return newValue
		}
		coords[x][y] = newValue
		lastValue = newValue
	}
}

func step(x, y int) (x1, y1 int) {
	// At the top?
	if y == 1 { // need a restart
		x1 = 1
		y1 = x + 1
		return x1, y1
	}
	// Move up and along
	x1 = x + 1
	y1 = y - 1
	return x1, y1
}
