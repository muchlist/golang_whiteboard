package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSet(t *testing.T) {
	setString := NewSet()
	success := setString.Add("Satu")
	success2 := setString.Add("Satu")

	sliceString := setString.Reveal()

	assert.Equal(t, true, success)
	assert.Equal(t, false, success2)
	assert.Equal(t, []string{"Satu"}, sliceString)

	setString.Add("Satu")
	setString.Add("Dua")
	setString.Add("Tiga")

	length := setString.Len()
	assert.Equal(t, 3, length)

	setString.Remove("Satu")
	assert.Equal(t, []string{"Dua", "Tiga"}, setString.Reveal())
}
