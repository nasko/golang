package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var qtyNoodles int
	var qtySauce float64

	for _, layer := range layers {
		switch layer {
		case "noodles":
			qtyNoodles += 50
		case "sauce":
			qtySauce += 0.2
		}
	}
	return qtyNoodles, qtySauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	if len(friendsList) < 1 {
		return
	}

	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantitiesForTwoPortions []float64, desiredPortionsCount int) []float64 {
	var desiredQuantities []float64

	for _, qtyForTwo := range quantitiesForTwoPortions {
		desiredQuantities = append(desiredQuantities, qtyForTwo*float64(desiredPortionsCount)/2)
	}
	return desiredQuantities
}
