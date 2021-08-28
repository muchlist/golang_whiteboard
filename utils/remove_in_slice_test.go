package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	strSlice := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	strSliceMod := RemoveAtIndex(strSlice, 4)
	assert.Equal(t, len(strSliceMod), len(strSlice)-1)
	assert.Equal(t, "sabtu", strSliceMod[4])
}

func TestRemoveAtIndexNotSorted(t *testing.T) {
	strSlice := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	strSliceMod := RemoveAtIndexNotSorted(strSlice, 4)
	assert.Equal(t, len(strSliceMod), len(strSlice)-1)
	assert.Equal(t, "minggu", strSliceMod[4])
}

// BenchmarkRemoveAtIndexNotSorted 1000000000	         0.04059 ns/op
func BenchmarkRemoveAtIndexNotSorted(b *testing.B) {
	strSlice := make([]string, 0)
	for i := 0; i < 100000; i++ {
		strSlice = append(strSlice, fmt.Sprintf("nilai %d", i))
	}
	for i := 0; i < 100000-1; i++ {
		RemoveAtIndexNotSorted(strSlice, i)
	}
}

// BenchmarkRemoveAtIndex 9546627600 ns/op
func BenchmarkRemoveAtIndex(b *testing.B) {
	strSlice := make([]string, 0)
	for i := 0; i < 100000; i++ {
		strSlice = append(strSlice, fmt.Sprintf("nilai %d", i))
	}
	for i := 0; i < 100000-1; i++ {
		RemoveAtIndex(strSlice, i)
	}
}
