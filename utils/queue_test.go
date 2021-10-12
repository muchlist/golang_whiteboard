package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("Satu")
	q.Enqueue("Dua")

	assert.Equal(t, 2, q.Size())

	antrian1, _ := q.Dequeue()
	front, _ := q.Front()
	antrian2, _ := q.Dequeue()
	antrian3, err := q.Dequeue()

	assert.Equal(t, "Satu", antrian1)
	assert.Equal(t, "Dua", front)
	assert.Equal(t, "Dua", antrian2)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, "", antrian3)
	assert.Error(t, err)
}
