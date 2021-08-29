package main

import "fmt"

func main() {
	text := "abcdEfg-!!!-efGhijk"
	encryptedText := encrypt(text, 1)
	fmt.Println(encryptedText)
}
