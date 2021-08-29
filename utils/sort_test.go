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

func TestSelectionSort(t *testing.T) {
	unsortedSlice := []int{2, 5, 6, 8, 9, 3, 5, 6, 8, 2, 4, 6, 2, 4, 7, 100, 80}
	fmt.Println(unsortedSlice)

	sortedList := SelectionSort(unsortedSlice)
	fmt.Println(sortedList)
	fmt.Println(unsortedSlice)
	assert.Equal(t, 100, sortedList[len(sortedList)-1])
	assert.Equal(t, 2, sortedList[0])
}

// BenchmarkBubbleSort-8   	  185316	      8764 ns/op
func BenchmarkBubbleSort(b *testing.B) {
	unsortedSlice := []int{2, 5, 6, 8, 9, 3, 5, 6, 8, 2, 4, 6, 2, 4, 7, 100, 80, 100, 1, 3, 2, 3, 6, 20, 21}
	unsortedSlice = append(unsortedSlice, unsortedSlice...)
	unsortedSlice = append(unsortedSlice, unsortedSlice...)
	for i := 0; i < b.N; i++ {
		_ = BubbleSort(unsortedSlice)
	}
}

// BenchmarkSelectionSort-8   	   90538	     11742 ns/op
func BenchmarkSelectionSort(b *testing.B) {
	unsortedSlice := []int{2, 5, 6, 8, 9, 3, 5, 6, 8, 2, 4, 6, 2, 4, 7, 100, 80, 100, 1, 3, 2, 3, 6, 20, 21}
	unsortedSlice = append(unsortedSlice, unsortedSlice...)
	unsortedSlice = append(unsortedSlice, unsortedSlice...)
	for i := 0; i < b.N; i++ {
		_ = SelectionSort(unsortedSlice)
	}
}
