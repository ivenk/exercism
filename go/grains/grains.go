package grains

import (
	"errors"
	"math"
)

func Square(field int) (value uint64, err error) {
	if field <= 0 || field > 64 {
		return 0, errors.New("Non valid input value")
	}

	// since we count the fields from 1 and not from 0 we subtracted one from the input
	field = field - 1
	result := uint64(math.Pow(float64(2), float64(field)))
	return result, nil
}

func Total() uint64 {
	result, _ := Square(64)
	result = result*2 - 1
	return result
}
