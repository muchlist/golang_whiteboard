package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStackString(t *testing.T) {
	stack := NewStackString()

	stack.Push("satu")
	stack.Push("dua")

	size := stack.Size()
	last, _ := stack.Pop()

	assert.Equal(t, 2, size)
	assert.Equal(t, "dua", last)
	size = stack.Size()

	assert.Equal(t, 1, size)
}
