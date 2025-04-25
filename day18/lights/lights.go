package lights

func New(data [][]int) *Lights {
	return &Lights{data}
}

type Lights struct {
	lights [][]int
}

func (ll *Lights) CountLights() int {
	total := 0
	for _, row := range ll.lights {
		for _, item := range row {
			if item == 1 {
				total++
			}
		}
	}

	return total
}

func (ll *Lights) SwitchLights() [][]int {
	output := deepCopy(ll.lights)

	for y := 0; y < len(ll.lights); y++ {
		for x := 0; x < len(ll.lights[y]); x++ {
			neightbourOnCount := ll.NeighbourOnCount(x, y)

			switch ll.lights[y][x] {
			case 1: // bulb on
				if !(neightbourOnCount == 2 || neightbourOnCount == 3) {
					output[y][x] = 0
				}
			case 0: // bulb off
				if neightbourOnCount == 3 {
					output[y][x] = 1
				}
			default:
				panic("not expected")
			}
		}
	}
	ll.lights = output // overwrite

	return ll.lights
}

func (ll *Lights) OverrideCorners() [][]int {
	output := deepCopy(ll.lights)

	ymax := len(ll.lights) - 1
	xmax := len(ll.lights[0]) - 1

	output[0][0] = 1
	output[ymax][0] = 1
	output[0][xmax] = 1
	output[ymax][xmax] = 1

	ll.lights = output // overwrite

	return ll.lights
}

func deepCopy(input [][]int) [][]int {
	output := [][]int{}
	for _, row := range input {
		row2 := make([]int, len(row))
		copy(row2, row)
		output = append(output, row2)
	}

	return output
}

func (ll *Lights) NeighbourOnCount(x, y int) int {
	onCount := 0

	ymax := len(ll.lights) - 1
	xmax := len(ll.lights[0]) - 1

	for yy := y - 1; yy < y+2; yy++ {
		for xx := x - 1; xx < x+2; xx++ {
			if xx == x && yy == y { // ignore current
				continue
			}
			if xx < 0 || yy < 0 || xx > xmax || yy > ymax { // edges
				continue
			}
			if ll.lights[yy][xx] == 1 {
				onCount++
			}
		}
	}

	return onCount
}
