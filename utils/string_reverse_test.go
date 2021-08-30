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

func TestReverseStringStack(t *testing.T) {
	str := "BAKWANA"
	strReversed := ReverseStringStack(str)

	assert.Equal(t, "ANAWKAB", strReversed)
}

// BenchmarkReverseString-8   	 3649749	       323.0 ns/op
func BenchmarkReverseString(b *testing.B) {
	for i:= 0; i < b.N ; i++ {
		str := "BAKWANAASDSASASASASASASASASASS"
		_ = ReverseString(str)
	}
}

// BenchmarkReverseStringStack-8   	  525902	      2267 ns/op
func BenchmarkReverseStringStack(b *testing.B) {
	for i:= 0; i < b.N ; i++ {
		str := "BAKWANAASDSASASASASASASASASASS"
		_ = ReverseStringStack(str)
	}
}
