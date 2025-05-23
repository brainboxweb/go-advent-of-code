package recipe

//revive:disable:max-control-nesting

func MixRatios(quant int) [][]int {
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

//revive:enable:max-control-nesting

func GetMaxScore(recipe Recipe, ratios [][]int, calorieCount int) int {
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

func NewIngredient(name string, capacity, durability, flavor, texture, calories int) Ingredient {
	return Ingredient{name, capacity, durability, flavor, texture, calories}
}

// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Recipe struct {
	RecipeIngredients []IngredientQuant
}

type IngredientQuant struct {
	ingredient Ingredient
	quantity   int
}

func (r *Recipe) AddIngredient(ingred Ingredient) {
	ri := IngredientQuant{ingred, 0}
	r.RecipeIngredients = append(r.RecipeIngredients, ri)
}

func (r *Recipe) SetRatios(ratios []int) {
	for i := 0; i < len(r.RecipeIngredients); i++ {
		r.RecipeIngredients[i].quantity = ratios[i]
	}
}

func (r *Recipe) GetScore(calorieCount int) int {
	// If calorie count not-zero, match or return
	if calorieCount != 0 {
		caloriesTotal := 0
		for _, ri := range r.RecipeIngredients {
			caloriesTotal += ri.ingredient.calories * ri.quantity
		}
		if caloriesTotal != calorieCount {
			return 0
		}
	}

	// capacity
	capacityTotal := 0
	for _, ri := range r.RecipeIngredients {
		capacityTotal += ri.ingredient.capacity * ri.quantity
	}
	if capacityTotal < 0 {
		capacityTotal = 0
	}

	// durability
	durabilityTotal := 0
	for _, ri := range r.RecipeIngredients {
		durabilityTotal += ri.ingredient.durability * ri.quantity
	}
	if durabilityTotal < 0 {
		durabilityTotal = 0
	}

	// flavor
	flavorTotal := 0
	for _, ri := range r.RecipeIngredients {
		flavorTotal += ri.ingredient.flavor * ri.quantity
	}
	if flavorTotal < 0 {
		flavorTotal = 0
	}

	// texture
	textureTotal := 0
	for _, ri := range r.RecipeIngredients {
		textureTotal += ri.ingredient.texture * ri.quantity
	}
	if textureTotal < 0 {
		textureTotal = 0
	}

	product := capacityTotal * durabilityTotal * flavorTotal * textureTotal

	return product
}
