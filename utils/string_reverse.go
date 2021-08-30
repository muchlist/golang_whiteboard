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

func ReverseStringStack(text string) string {
	stack := NewStackString()
	runeStr := make([]rune, len(text))
	for _, char := range text {
		stack.Push(string(char))
	}

	for i:= 0; i < len(text); i++ {
		r, _ := stack.Pop()
		runeStr[i] = []rune(r)[0]
	}

	return string(runeStr)
}