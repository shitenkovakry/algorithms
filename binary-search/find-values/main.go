package main

import (
	"math"
)

const (
	minimumNumbersOfValuesInAnArray = 1
)

func FindTheMaximumNumberOfCheks(array []int) int {
	numberOfValues := len(array)

	if numberOfValues < minimumNumbersOfValuesInAnArray {

		return 0
	}

	maximumNumberOfCheks := math.Log2(float64(numberOfValues))

	return int(maximumNumberOfCheks)
}
