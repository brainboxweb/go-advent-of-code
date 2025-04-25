package recipe_test

import (
	"testing"

	"github.com/brainboxweb/advent/day15/recipe"
	"github.com/stretchr/testify/require"
)

func TestScore(t *testing.T) {
	butterscotch := recipe.NewIngredient("Butterscotch", -1, -2, 6, 3, 8)
	cinnamon := recipe.NewIngredient("Cinnamon", 2, 3, -2, -1, 3)

	r := recipe.Recipe{}
	r.AddIngredient(butterscotch)
	r.AddIngredient(cinnamon)

	r.SetRatios([]int{44, 56})

	expected := 62842880
	result := r.GetScore(0)
	require.Equal(t, expected, result)
}

func TestScoreMixing(t *testing.T) {
	butterscotch := recipe.NewIngredient("Butterscotch", -1, -2, 6, 3, 8)
	cinnamon := recipe.NewIngredient("Cinnamon", 2, 3, -2, -1, 3)

	r := recipe.Recipe{}
	r.AddIngredient(butterscotch)
	r.AddIngredient(cinnamon)

	ratios := recipe.MixRatios(len(r.RecipeIngredients))

	expected := 62842880
	result := recipe.GetMaxScore(r, ratios, 0)
	require.Equal(t, expected, result)
}
