package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
*/
var test = []struct {
	input    string
	expected Ingredient
}{
	{
		"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
		Ingredient{"Butterscotch", -1, -2, 6, 3, 8},
	},
}

func TestParse(t *testing.T) {
	for _, test := range test {
		if actual := parse(test.input); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

/*
Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of cinnamon (because the amounts of each

	ingredient must add up to 100) would result in a cookie with the following properties:

A capacity of 44*-1 + 56*2 = 68
A durability of 44*-2 + 56*3 = 80
A flavor of 44*6 + 56*-2 = 152
A texture of 44*3 + 56*-1 = 76
*/
func TestScore(t *testing.T) {

	Butterscotch := Ingredient{"Butterscotch", -1, -2, 6, 3, 8}
	Cinnamon := Ingredient{"Cinnamon", 2, 3, -2, -1, 3}

	recipe := Recipe{}
	recipe.AddIngredient(Butterscotch)
	recipe.AddIngredient(Cinnamon)

	recipe.SetRatios([]int{44, 56})

	expected := 62842880
	result := recipe.GetScore(0)
	require.Equal(t, expected, result)

}

func TestScoreMixing(t *testing.T) {

	Butterscotch := Ingredient{"Butterscotch", -1, -2, 6, 3, 8}
	Cinnamon := Ingredient{"Cinnamon", 2, 3, -2, -1, 3}

	recipe := Recipe{}
	recipe.AddIngredient(Butterscotch)
	recipe.AddIngredient(Cinnamon)

	ratios := mixRatios(len(recipe.recipeIngredients))

	expected := 62842880
	result := getMaxScore(recipe, ratios, 0)
	require.Equal(t, expected, result)

}

var test2 = []struct {
	input    string
	expected int
}{
	{
		`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
		62842880,
	},
	{
		day15data,
		18965440,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		if actual := run(test.input, 0); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

var test3 = []struct {
	input    string
	expected int
}{
	{
		`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
		57600000,
	},
	{
		day15data,
		15862900,
	},
}

func TestRun500(t *testing.T) {
	for _, test := range test3 {
		if actual := run(test.input, 500); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}

const day15data = `Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1`
