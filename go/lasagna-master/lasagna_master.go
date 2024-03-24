package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	est_time := 2
	if time != 0 {
		est_time = time
	}
	return len(layers) * est_time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, e := range layers {
		if e == "sauce" {
			sauce += 0.2
		} else if e == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var res []float64
	scale := float64(portion) / 2
	for _, e := range quantities {
		res = append(res, e*scale)
	}
	return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
