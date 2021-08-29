package main

import (
	"fmt"
	"github.com/muchlist/golang_whiteboard/utils"
)

func main() {
	text := "TENET"
	if isPalindrome1(text) {
		fmt.Println("Text adalah palindrome")
	} else {
		fmt.Println("Text tidak palindrome")
	}

	if isPalindrome2(text) {
		fmt.Println("Text adalah palindrome")
	} else {
		fmt.Println("Text tidak palindrome")
	}
}

// melakukan perulangan sebanyak setengah text
func isPalindrome1(text string) bool {
	runeText := []rune(text)
	midIndex := len(runeText)
	for i, j := 0, len(runeText)-1; i < midIndex; i, j = i+1, j-1 {
		if runeText[i] != runeText[j] {
			return false
		}
	}
	return true
}

// membalik text secara keseluruhan lalu membandingkannya
func isPalindrome2(text string) bool {
	return text == utils.ReverseString(text)
}
