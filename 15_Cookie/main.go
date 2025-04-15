package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func run(input string, calorieCount int) int {
	b := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(b)

	recipe := Recipe{}

	for scanner.Scan() {
		text := scanner.Text()
		ingred := parse(text)
		recipe.AddIngredient(ingred)
	}

	ratios := mixRatios(len(recipe.recipeIngredients))

	return getMaxScore(recipe, ratios, calorieCount)
}

func mixRatios(quant int) [][]int {

	ratios := [][]int{}

	if quant == 2 {
		for i := 0; i < 100; i++ {
			ratio := []int{i, 100 - i}
			ratios = append(ratios, ratio)
		}
	}

	if quant == 4 {
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				for k := 0; k < 100; k++ {
					for l := 0; l < 100; l++ {
						if (i + j + k + l) == 100 {
							ratio := []int{i, j, k, l}
							ratios = append(ratios, ratio)
						}
					}
				}
			}
		}
	}

	return ratios
}

func getMaxScore(recipe Recipe, ratios [][]int, calorieCount int) int {
	maxScore := 0
	for _, ratio := range ratios {
		recipe.SetRatios(ratio)
		result := recipe.GetScore(calorieCount)
		if result > maxScore {
			maxScore = result
		}
	}
	return maxScore
}

//Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
func parse(phrase string) Ingredient {
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

	ingred := Ingredient{name, capacity, durability, flavor, texture, calories}

	return ingred
}

//Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Recipe struct {
	recipeIngredients []RecipeIngredient
}

type RecipeIngredient struct {
	ingredient Ingredient
	quantity   int
}

func (r *Recipe) AddIngredient(ingred Ingredient) {
	ri := RecipeIngredient{ingred, 0}
	r.recipeIngredients = append(r.recipeIngredients, ri)
}

func (r *Recipe) SetRatios(ratios []int) {
	for i := 0; i < len(r.recipeIngredients); i++ {
		r.recipeIngredients[i].quantity = ratios[i]
	}
}

func (r *Recipe) GetScore(calorieCount int) int {

	//If calorie count not-zero, match or return
	if calorieCount != 0 {
		caloriesTotal := 0
		for _, ri := range r.recipeIngredients {
			caloriesTotal += ri.ingredient.calories * ri.quantity
		}
		if caloriesTotal != calorieCount {
			return 0
		}
	}

	//capacity
	capacityTotal := 0
	for _, ri := range r.recipeIngredients {
		capacityTotal += ri.ingredient.capacity * ri.quantity
	}
	if capacityTotal < 0 {
		capacityTotal = 0
	}

	//durability
	durabilityTotal := 0
	for _, ri := range r.recipeIngredients {
		durabilityTotal += ri.ingredient.durability * ri.quantity
	}
	if durabilityTotal < 0 {
		durabilityTotal = 0
	}

	//flavor
	flavorTotal := 0
	for _, ri := range r.recipeIngredients {
		flavorTotal += ri.ingredient.flavor * ri.quantity

	}
	if flavorTotal < 0 {
		flavorTotal = 0
	}

	//texture
	textureTotal := 0
	for _, ri := range r.recipeIngredients {
		textureTotal += ri.ingredient.texture * ri.quantity
	}
	if textureTotal < 0 {
		textureTotal = 0
	}

	product := capacityTotal * durabilityTotal * flavorTotal * textureTotal

	return product
}
