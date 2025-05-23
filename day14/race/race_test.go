package race_test

import (
	"testing"

	"github.com/brainboxweb/advent/day14/race"
	"github.com/stretchr/testify/require"
)

func TestGetDistance(t *testing.T) {
	theReindeer := race.NewReindeer("Comet", 14, 10, 127)
	time := 10
	expected := 140
	result := theReindeer.GetDistance(time)
	require.Equal(t, expected, result)

	theReindeer2 := race.NewReindeer("Comet", 14, 10, 127)
	time2 := 1000
	expected2 := 1120
	result2 := theReindeer2.GetDistance(time2)
	require.Equal(t, expected2, result2)

	theReindeer3 := race.NewReindeer("Dancer", 16, 11, 162)
	time3 := 1000
	expected3 := 1056
	result3 := theReindeer3.GetDistance(time3)
	require.Equal(t, expected3, result3)
}

/*
Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in
the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets
his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated
139 points by the 140th second.
*/
func TestGetScore(t *testing.T) {
	comet := race.NewReindeer("Comet", 14, 10, 127)
	dancer := race.NewReindeer("Dancer", 16, 11, 162)

	r := race.Race{}
	r.AddReindeer(comet)
	r.AddReindeer(dancer)

	r.RunRace(140)

	topScore := r.GetTopScore()
	require.Equal(t, 139, topScore)
}

/*
After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with
the new scoring system, Dancer would win (if the race ended at 1000 seconds).
*/
func TestGetScore2(t *testing.T) {
	comet := race.NewReindeer("Comet", 14, 10, 127)
	dancer := race.NewReindeer("Dancer", 16, 11, 162)

	r := race.Race{}
	r.AddReindeer(comet)
	r.AddReindeer(dancer)

	r.RunRace(1000)

	topScore := r.GetTopScore()
	require.Equal(t, 689, topScore)
}
