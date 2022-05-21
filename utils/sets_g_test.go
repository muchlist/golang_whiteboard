package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSetG(t *testing.T) {
	setT := NewSetG[int]()
	success := setT.Add(1)
	success2 := setT.Add(1)

	sliceT := setT.Reveal()

	assert.Equal(t, true, success)
	assert.Equal(t, false, success2)
	assert.Equal(t, []int{1}, sliceT)

	setT.Add(1)
	setT.Add(2)
	setT.Add(3)

	length := setT.Len()
	assert.Equal(t, 3, length)

	setT.Remove(1)
	assert.Equal(t, []int{2, 3}, setT.Reveal())
}
