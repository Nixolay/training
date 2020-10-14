// Package thirt ...
package thirt

import (
	"strconv"
)

// Thirt ...
func Thirt(n int) int {
	// temporare result that is storedand compared with each iteration
	tempPrevResult := recursiveCalculation(n)

	// iterating until condition met
	for result := n; result != tempPrevResult; result = recursiveCalculation(result) {
		tempPrevResult = result
	}

	return tempPrevResult
}

// function that is used for recursive calls.
func recursiveCalculation(n int) int {
	numbers := [6]int{1, 10, 9, 12, 3, 4}
	numberAsString := strconv.Itoa(n)
	numbersIndex := 0
	result := 0

	for i := len(numberAsString); i > 0; i-- {
		singleInt, _ := strconv.Atoi(string(numberAsString[i-1]))
		result += singleInt * numbers[numbersIndex%6]
		numbersIndex++
	}

	return result
}
