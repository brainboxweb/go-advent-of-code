package main

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

/*
Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
*/
var test = []struct {
	input    string
	expected []string
}{
	{
		"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
		[]string{"Comet", "14", "10", "127"},
	},
	{
		"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
		[]string{"Dancer", "16", "11", "162"},
	},
}

func TestParse(t *testing.T) {
	for _, test := range test {
		if actual := parse(test.input); !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

func TestGetDistance(t *testing.T) {

	theReindeer := Reindeer{name: "Comet", speed: 14, flyTime: 10, restTime: 127}
	time := 10
	expected := 140
	result := theReindeer.getDistance(time)
	require.Equal(t, expected, result)

	theReindeer2 := Reindeer{name: "Comet", speed: 14, flyTime: 10, restTime: 127}
	time2 := 1000
	expected2 := 1120
	result2 := theReindeer2.getDistance(time2)
	require.Equal(t, expected2, result2)

	theReindeer3 := Reindeer{name: "Dancer", speed: 16, flyTime: 11, restTime: 162}
	time3 := 1000
	expected3 := 1056
	result3 := theReindeer3.getDistance(time3)
	require.Equal(t, expected3, result3)
}

var test2 = []struct {
	input    string
	time     int
	expected int
}{
	{
		`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
		1000,
		1120,
	},
	{
		day14data,
		2503,
		2655,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		if actual := run(test.input, test.time); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

/*
Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in
the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets
his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated
139 points by the 140th second.
*/
func TestGetScore(t *testing.T) {

	comet := Reindeer{name: "Comet", speed: 14, flyTime: 10, restTime: 127}
	dancer := Reindeer{name: "Dancer", speed: 16, flyTime: 11, restTime: 162}

	race := Race{}
	race.addReindeer(&comet)
	race.addReindeer(&dancer)

	race.runRace(140)

	topScore := race.getTopScore()
	require.Equal(t, 139, topScore)
}

/*
After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with
the new scoring system, Dancer would win (if the race ended at 1000 seconds).
*/
func TestGetScore2(t *testing.T) {

	comet := Reindeer{name: "Comet", speed: 14, flyTime: 10, restTime: 127}
	dancer := Reindeer{name: "Dancer", speed: 16, flyTime: 11, restTime: 162}

	race := Race{}
	race.addReindeer(&comet)
	race.addReindeer(&dancer)

	race.runRace(1000)

	topScore := race.getTopScore()
	require.Equal(t, 689, topScore)
}

var test3 = []struct {
	input    string
	time     int
	expected int
}{
	{
		`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
		1000,
		689,
	},
	{
		day14data,
		2503,
		1059,
	},
}

func TestRun2(t *testing.T) {
	for _, test := range test3 {
		if actual := run2(test.input, test.time); actual != test.expected {
			t.Errorf("Convert(%s) = %d, expected %d.",
				test.input, actual, test.expected)
		}
	}
}

const day14data = `Vixen can fly 8 km/s for 8 seconds, but then must rest for 53 seconds.
Blitzen can fly 13 km/s for 4 seconds, but then must rest for 49 seconds.
Rudolph can fly 20 km/s for 7 seconds, but then must rest for 132 seconds.
Cupid can fly 12 km/s for 4 seconds, but then must rest for 43 seconds.
Donner can fly 9 km/s for 5 seconds, but then must rest for 38 seconds.
Dasher can fly 10 km/s for 4 seconds, but then must rest for 37 seconds.
Comet can fly 3 km/s for 37 seconds, but then must rest for 76 seconds.
Prancer can fly 9 km/s for 12 seconds, but then must rest for 97 seconds.
Dancer can fly 37 km/s for 1 seconds, but then must rest for 36 seconds.`
