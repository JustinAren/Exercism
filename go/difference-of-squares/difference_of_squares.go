package diffsquares

func SquareOfSum(n int) int {
	sum := (1 + float64(n)) / 2 * float64(n)
	return int(sum * sum)
}

func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
	if difference < 0 {
		return difference * -1
	}
	return difference
}
