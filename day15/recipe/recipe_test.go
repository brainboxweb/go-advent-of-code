package recipe_test

import (
	"testing"

	"github.com/brainboxweb/advent/day15/recipe"
	"github.com/stretchr/testify/require"
)

func TestScore(t *testing.T) {
	Butterscotch := recipe.NewIngredient("Butterscotch", -1, -2, 6, 3, 8)
	Cinnamon := recipe.NewIngredient("Cinnamon", 2, 3, -2, -1, 3)

	recipe := recipe.Recipe{}
	recipe.AddIngredient(Butterscotch)
	recipe.AddIngredient(Cinnamon)

	recipe.SetRatios([]int{44, 56})

	expected := 62842880
	result := recipe.GetScore(0)
	require.Equal(t, expected, result)
}

func TestScoreMixing(t *testing.T) {
	Butterscotch := recipe.NewIngredient("Butterscotch", -1, -2, 6, 3, 8)
	Cinnamon := recipe.NewIngredient("Cinnamon", 2, 3, -2, -1, 3)

	r := recipe.Recipe{}
	r.AddIngredient(Butterscotch)
	r.AddIngredient(Cinnamon)

	ratios := recipe.MixRatios(len(r.RecipeIngredients))

	expected := 62842880
	result := recipe.GetMaxScore(r, ratios, 0)
	require.Equal(t, expected, result)
}
