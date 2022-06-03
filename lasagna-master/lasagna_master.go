package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, e := range layers {
		switch e {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numOfPortions int) []float64 {
	scaledQuantities := []float64{}
	for _, e := range amounts {
		scaledQuantities = append(scaledQuantities, (e/2)*float64(numOfPortions))
	}
	return scaledQuantities
}
