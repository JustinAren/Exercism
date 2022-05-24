package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("number cannot be smaller than 1")
	}
	if number > 64 {
		return 0, errors.New("number cannot be bigger than 64")
	}
	return uint64(math.Pow(float64(2), float64(number-1))), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		square, _ := Square(i)
		total += square
	}
	return total
}
