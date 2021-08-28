package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	str := "BAKWANA"
	strReversed := ReverseString(str)

	assert.Equal(t, "ANAWKAB", strReversed)
}
