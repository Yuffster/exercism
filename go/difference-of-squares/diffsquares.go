package diffsquares

import "math"

const testVersion = 2

func SquareOfSums(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2))
}

func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return total
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
