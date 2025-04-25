package race

import (
	"sort"
)

func NewReindeer(name string, speed, flyTime, restTime int) *Reindeer {
	return &Reindeer{name: name, speed: speed, flyTime: flyTime, restTime: restTime}
}

type Reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
}

func (r *Reindeer) GetDistance(time int) int {
	totalFlyTime := 0

	// Loops
	loopTime := r.flyTime + r.restTime
	chunks := time / loopTime // Taking advantance of integer maths
	totalFlyTime += chunks * r.flyTime

	// Partials
	remainderTime := time - (chunks * loopTime)
	if remainderTime > r.flyTime {
		totalFlyTime += r.flyTime
	} else {
		totalFlyTime += remainderTime
	}

	return totalFlyTime * r.speed
}

type Race struct {
	scores map[*Reindeer]int
}

func (race *Race) AddReindeer(reindeer *Reindeer) {
	// Avoid nil map panic
	if race.scores == nil {
		race.scores = make(map[*Reindeer]int)
	}
	race.scores[reindeer] = 0
}

func (race *Race) RunRace(time int) {
	// Award points every second
	for i := 1; i < time+1; i++ {
		dists := distances{}

		for rein := range race.scores {
			dist := distance{rein, rein.GetDistance(i)}
			dists = append(dists, dist)
		}

		// Dists still has reference to the real reindeer
		sort.Sort(dists)

		winningDistance := dists[0].Distance

		for i := 0; i < len(dists); i++ {
			if winningDistance == dists[i].Distance {
				theReindeer := dists[i].Reindeer
				race.scores[theReindeer]++
			}
		}
	}
}

func (race *Race) GetTopScore() int {
	topScore := 0
	for _, score := range race.scores {
		if score > topScore {
			topScore = score
		}
	}
	return topScore
}

// ------- sorting
type distance struct {
	Reindeer *Reindeer
	Distance int
}

type distances []distance

func (d distances) Len() int      { return len(d) }
func (d distances) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d distances) Less(i, j int) bool {
	return d[i].Distance > d[j].Distance
}
