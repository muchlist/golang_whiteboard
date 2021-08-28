package utils

func ReverseString(text string) string {
	if len(text) < 2 {
		return text
	}
	runeStr := []rune(text)
	midIndex := len(runeStr) / 2
	for i, j := 0, len(runeStr)-1; i < midIndex; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}

	return string(runeStr)
}
