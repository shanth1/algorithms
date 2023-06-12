package algorithms

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middleIndex := len(arr) / 2
	pivot := arr[middleIndex]

	var less []int
	var medium = []int{pivot}
	var greater []int

	for _, el := range append(arr[:middleIndex], arr[middleIndex+1:]...) {
		if el < pivot {
			less = append(less, el)
		} else if el > pivot {
			greater = append(greater, el)
		} else {
			medium = append(medium, el)
		}
	}

	return append(append(QuickSort(less), medium...), QuickSort(greater)...)
}
