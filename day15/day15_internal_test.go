package day15

import (
	"testing"

	"github.com/brainboxweb/advent/day15/recipe"
)

func TestParse(t *testing.T) {
	var test = []struct {
		input    string
		expected recipe.Ingredient
	}{
		{
			"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
			recipe.NewIngredient("Butterscotch", -1, -2, 6, 3, 8),
		},
	}
	for _, test := range test {
		if actual := parse(test.input); actual != test.expected {
			t.Errorf("Parse(%s) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
}
