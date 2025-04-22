package day15

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/brainboxweb/advent/day15/recipe"
)

func Part1(input string, calorieCount int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	myRecipe := recipe.Recipe{}

	for scanner.Scan() {
		text := scanner.Text()
		ingred := parse(text)
		myRecipe.AddIngredient(ingred)
	}

	ratios := recipe.MixRatios(len(myRecipe.RecipeIngredients))

	return recipe.GetMaxScore(myRecipe, ratios, calorieCount)
}

// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
func parse(phrase string) recipe.Ingredient {
	phrase = strings.Trim(phrase, ".")
	tokens := strings.Split(phrase, " ")

	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.Trim(tokens[i], ":")
		tokens[i] = strings.Trim(tokens[i], ",")
	}

	name := strings.Trim(tokens[0], ":")
	capacity, _ := strconv.Atoi(tokens[2])
	durability, _ := strconv.Atoi(tokens[4])
	flavor, _ := strconv.Atoi(tokens[6])
	texture, _ := strconv.Atoi(tokens[8])
	calories, _ := strconv.Atoi(tokens[10])

	ingredient := recipe.NewIngredient(name, capacity, durability, flavor, texture, calories)

	return ingredient
}
