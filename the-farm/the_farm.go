package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
var SillyNephewError = errors.New("")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// panic("Please implement DivideFood")
	fodderReturn, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodderReturn >= 0 {
		return fodderReturn * 2 / float64(cows), err
	} else if err != ErrScaleMalfunction && fodderReturn > 0 {
		return 0, err
	} else if fodderReturn <= 0 && err == ErrScaleMalfunction || err == nil {
		return 0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	}

}
