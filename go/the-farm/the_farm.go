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
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	errNegativeFodder := errors.New("negative fodder")
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, &SillyNephewError{
			cows: cows,
		}
	}

	if err == ErrScaleMalfunction {
		if fodder > 0 {
			return fodder * 2 / float64(cows), nil
		}

		return 0, errNegativeFodder
	}

	if err != nil {
		return 0, err
	}

	if fodder < 0 {
		return 0, errNegativeFodder
	}

	return fodder / float64(cows), nil
}
