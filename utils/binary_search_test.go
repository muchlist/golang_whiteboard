package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinSe(t *testing.T) {
	arrayX := []int{1, 2, 3, 4, 5, 6, 10, 12, 18}

	index, err := BinarySearch(arrayX, 3)
	assert.Equal(t, 2, index)
	assert.Nil(t, err)
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayX := []int{1, 2, 3, 4, 5, 6, 10, 12, 18}
		_, _ = BinarySearch(arrayX, 3)
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayX := []int{1, 2, 3, 4, 5, 6, 10, 12, 18}
		_, _ = BinarySearchRecursive(arrayX, 3, 0, len(arrayX)-1)
	}
}
