package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	countOfSauce, countOfNoodles := 0, 0
	for _, v := range layers {
		switch v {
		case "sauce":
			countOfSauce++
		case "noodles":
			countOfNoodles++
		}
	}
	return countOfNoodles * 50, float64(countOfSauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) []string {
	myList[len(myList)-1] = friendList[len(friendList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numOfPortion int) []float64 {
	scaledRecipe := []float64{}
	for _, v := range quantities {
		scaledRecipe = append(scaledRecipe, v/2*float64(numOfPortion))
	}
	return scaledRecipe
}
