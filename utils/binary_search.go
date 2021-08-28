package utils

import "errors"

// BinarySearch return index from sorted slice
// sortedSlice must be sorted list
// if index not found return -1 and error
func BinarySearch(sortedSlice []int, value int) (int, error) {
	low := 0
	high := len(sortedSlice) - 1

	for low <= high {
		middle := (low + high) / 2
		if value == sortedSlice[middle] {
			return middle, nil
		} else if value < sortedSlice[middle] {
			high = middle - 1
		} else { // value > sortedSlice[middle]
			low = middle + 1
		}
	}
	return -1, errors.New("index not found")
}

// BinarySearchRecursive return index from sorted slice
// sortedSlice must be sorted list
// low = 0, and high is len(sortedSlice) - 1
// if index not found return -1 and error
func BinarySearchRecursive(sortedSlice []int, value, low, high int) (index int, err error) {
	middle := (low + high) / 2

	switch {
	case low > high: // BREAKER
		index = -1
		err = errors.New("index not found")
	case sortedSlice[middle] == value:
		index = middle
		err = nil
	case sortedSlice[middle] > value:
		index, err = BinarySearchRecursive(sortedSlice, value, low, middle)
	default:
		index, err = BinarySearchRecursive(sortedSlice, value, middle+1, high)
	}

	return
}
