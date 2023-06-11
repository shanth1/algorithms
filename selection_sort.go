package algorithms

func findSmallest(arr []int) (element, index int) {
	element = arr[0]
	index = 0
	for i, v := range arr {
		if v < element {
			element = v
			index = i
		}
	}
	return element, index
}

func SelectionSort(arr []int) []int {
	var newArr []int
	for range arr {
		smallest, smallestIndex := findSmallest(arr)
		arr = append(arr[:smallestIndex], arr[(smallestIndex+1):]...)
		newArr = append(newArr, smallest)
	}
	return newArr
}
