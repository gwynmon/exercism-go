package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, averageTime int) int {
    if averageTime == 0 {
        averageTime = 2
    }
    return len(slice) * averageTime
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (int, float64) {
	var noodles int = 0
    var sauce float64 = 0.0
    for _, element := range slice {
        if(element == "noodles") {
            noodles += 50
        } else if(element == "sauce") {
            sauce += 0.2
        }
    }
    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := make([]float64, len(quantities))
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * float64(portions)/2.0
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
