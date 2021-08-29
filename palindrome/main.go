package main

import "fmt"

func main() {
	text := "TENNET"
	if isPalindrome1(text) {
		fmt.Println("Text adalah palindrome")
	} else {
		fmt.Println("Text tidak palindrome")
	}
}

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
