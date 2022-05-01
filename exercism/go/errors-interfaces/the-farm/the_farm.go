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
func DivideFood(weightFodder WeightFodder, cowsCnt int) (float64, error) {
	wfFodder, wfError := weightFodder.FodderAmount()

	switch {
	case wfError == ErrScaleMalfunction:
		if wfFodder >= 0 {
			return wfFodder / float64(cowsCnt) * 2, nil
		} else {
			return 0.0, errors.New("negative fodder")
		}
	case wfError != nil:
		return 0.0, wfError
	case cowsCnt == 0:
		return 0.0, errors.New("division by zero")
	case wfFodder < 0:
		return 0.0, errors.New("negative fodder")
	case cowsCnt < 0:
		return 0.0, &SillyNephewError{
			cows: cowsCnt,
		}
	}
	return wfFodder / float64(cowsCnt), nil
}
