package algorithms

func Factorial(a int) int {
	if a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}
