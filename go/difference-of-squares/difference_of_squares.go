package diffsquares

import (
	"math"
)

func SquareOfSum(input int) int {
	var sum int
	for i := 1; i <= input; i++ {
		sum += i
	}
	square := math.Pow(float64(sum), 2)
	return int(square)
}

func SumOfSquares(input int) int {
	var sum int
	for i := 1; i <= input; i++ {
		square := math.Pow(float64(i), 2)
		sum += int(square)
	}
	return sum
}

func Difference(input int) int {
	difference := SquareOfSum(input) - SumOfSquares(input)
	return difference
}
