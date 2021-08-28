package utils

func RemoveAtIndex(sliceT []string, index int) []string {
	return append(sliceT[:index], sliceT[index+1:]...)
}

func RemoveAtIndexNotSorted(sliceT []string, index int) []string {
	sliceT[index] = sliceT[len(sliceT)-1]
	return sliceT[:len(sliceT)-1]
}
