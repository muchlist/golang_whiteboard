package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChipper(t *testing.T) {
	text := "abcdEfg-!"
	encryptedText := encrypt(text, 1)
	encryptedText2 := caesarCipher(text, 1)

	assert.Equal(t, "bcdeFgh-!", encryptedText)
	assert.Equal(t, "bcdeFgh-!", encryptedText2)
}

// BenchmarkChipper1-8   	 2765098	       541.6 ns/op
func BenchmarkChipper1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		text := "abcdEfg-!asdasdadasdasdadsad"
		_ = encrypt(text, 12)
	}
}

// BenchmarkChipper2-8   	  272872	      4813 ns/op
func BenchmarkChipper2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		text := "abcdEfg-!asdasdadasdasdadsad"
		_ = caesarCipher(text, 12)
	}
}
