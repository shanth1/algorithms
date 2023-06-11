package algorithms

func Factorial(a int) int {
	if a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + Sum(arr[1:])
}
