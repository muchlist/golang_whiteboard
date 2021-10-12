package main

import "fmt"

func main() {
	text := "abcdEfg-!!!-efGhijk"
	encryptedText := caesarChipperImprove(text, 5)
	fmt.Println(encryptedText)
}
