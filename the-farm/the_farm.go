package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// panic("Please implement DivideFood")
	fodderReturn, err := weightFodder.FodderAmount()
	switch {
	case err == ErrScaleMalfunction && fodderReturn > 0:
		return fodderReturn * 2 / float64(cows), nil
	case err != nil && fodderReturn > 0:
		return 0, err
	case (err == nil || err == ErrScaleMalfunction) && fodderReturn < 0:
		return 0, errors.New("negative fodder")
	case (err != nil && err != ErrScaleMalfunction) && fodderReturn < 0:
		return 0, err
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{
			cows: cows,
		}
	}

	return fodderReturn / float64(cows), nil
}
