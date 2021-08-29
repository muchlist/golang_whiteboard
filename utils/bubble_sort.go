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
