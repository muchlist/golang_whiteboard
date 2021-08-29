package utils

func BubbleSort(sliceX []int) []int {
	copySlice := make([]int, len(sliceX))
	copy(copySlice, sliceX)

	unsortedUntilIndex := len(copySlice) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			if copySlice[i] > copySlice[i+1] {
				copySlice[i], copySlice[i+1] = copySlice[i+1], copySlice[i]
				sorted = false
			}
		}
		unsortedUntilIndex--
	}
	return copySlice
}

func SelectionSort(arrayX []int) []int {
	array := make([]int, len(arrayX))
	copy(array, arrayX)

	// len(array) - 1 karena tidak perlu melakukan perulangan untuk yang terakhir
	for i := 0; i < len(array)-1; i++ {
		lowestNumberIndex := i
		// mencari index dengan value terkecil
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}
		if lowestNumberIndex != i {
			// jika array i bukan yang terkecil tukar dengan lowerNumberIndex
			array[i], array[lowestNumberIndex] = array[lowestNumberIndex], array[i]
		}
	}
	return array
}
