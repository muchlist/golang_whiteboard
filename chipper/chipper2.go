package main

import "strings"

func caesarCipher(s string, k int32) string {
	// Write your code here
	lowerAlp := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"} // index 0 - 25

	// convert string to slice rune
	// runs := []rune(s)
	var strModified strings.Builder
	for _, char := range s {
		isCharUpperCase := isUpperCase(string(char))
		indexChar := findIndexInAlphabet(string(char), lowerAlp[:])
		if indexChar == -1 {
			// bukan alphabet, return seaslinya
			strModified.WriteString(string(char))
			continue
		}
		// untuk setiap huruf, cari index keberapa lalu tambahkan dengan k % 26
		indexModified := (indexChar + int(k)) % len(lowerAlp)
		if isCharUpperCase {
			strModified.WriteString(strings.ToUpper(lowerAlp[indexModified]))
		} else {
			strModified.WriteString(lowerAlp[indexModified])
		}
	}
	return strModified.String()
}

func findIndexInAlphabet(s string, lowerAlp []string) int {
	for i, char := range lowerAlp {
		if strings.ToLower(s) == char {
			return i
		}
	}
	return -1
}

func isUpperCase(s string) bool {
	return strings.ToUpper(s) == s
}
