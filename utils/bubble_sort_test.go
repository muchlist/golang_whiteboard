package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	unsortedSlice := []int{2, 5, 6, 8, 9, 3, 5, 6, 8, 2, 4, 6, 2, 4, 7, 100, 80}
	fmt.Println(unsortedSlice)

	sortedList := BubbleSort(unsortedSlice)
	fmt.Println(sortedList)
	fmt.Println(unsortedSlice)
	assert.Equal(t, 100, sortedList[len(sortedList)-1])
	assert.Equal(t, 2, sortedList[0])
}
