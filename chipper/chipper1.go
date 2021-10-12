package main

func caesarChipperImprove(text string, inc int32) string {
	runeText := []rune(text)
	for i, char := range runeText {
		runeText[i] = chipper(char, inc)
	}
	return string(runeText)
}

func chipper(char rune, inc int32) rune {

	inc = inc % 26

	// selain a-z A-Z return seadanya
	if char > 'z' || char < 'a' && char > 'Z' || char < 'A' {
		return char
	}

	// huruf kecil
	if char <= 'z' && char >= 'a' {
		// melakukan pergeseran huruf
		char = char + inc
		if char > 'z' {
			char = char - 26
		}
		if char < 'a' {
			char = char + 26
		}
		return char
	} else {
		// huruf besar
		// melakukan pergeseran huruf
		char = char + inc
		if char > 'Z' {
			char = char - 26
		}
		if char < 'A' {
			char = char + 26
		}
		return char
	}
}
