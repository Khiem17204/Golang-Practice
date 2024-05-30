package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood calculates the amount of food per cow based on the provided FodderCalculator
// and the number of cows. It returns the amount of food per cow as a float64 or an error.
func DivideFood(fd FodderCalculator, cows int) (float64, error) {
	// Retrieve the total amount of fodder for all the cows
	totalFodder, err := fd.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	// Get the fattening factor
	fatteningFactor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}

	// Calculate the amount of food per cow
	foodPerCow := totalFodder * fatteningFactor / float64(cows)
	return foodPerCow, nil
}

// ValidateInputAndDivideFood validates the number of cows and then divides the food
// based on the provided FodderCalculator. It returns the amount of food per cow as a float64 or an error.
func ValidateInputAndDivideFood(fd FodderCalculator, cows int) (float64, error) {
	// Validate the number of cows
	if cows < 1 {
		return 0.0, errors.New("invalid number of cows")
	}
	// If the number of cows is valid, divide the food
	return DivideFood(fd, cows)
}

// InvalidCows is a custom error type for invalid number of cows.
type InvalidCows struct {
	NumCows int
	message string
}

// Error implements the error interface for InvalidCows.
func (e *InvalidCows) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumCows, e.message)
}

// ValidateNumberOfCows validates the number of cows and returns an error if invalid.
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCows{NumCows: cows, message: "there are no negative cows"}
	case cows == 0:
		return &InvalidCows{NumCows: cows, message: "no cows don't need food"}
	default:
		return nil
	}
}
