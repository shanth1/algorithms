package algorithms

func Binary_search(list []int, item int) (index int, ok bool) {
	low := 0
	high := len(list) - 1

	var mid int
	var guess int

	for low <= high {
		mid = (low + high) / 2
		guess = list[mid]

		if guess == item {
			return mid, true
		}
		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, false
}
