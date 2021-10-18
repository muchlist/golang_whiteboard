package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := int64(1)
	result, err := flippingBits(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func flippingBits(n int64) (int64, error) {
	// ubah int64 kedalam string dengan format 32 bit
	str32Bit := fmt.Sprintf("%.32b", n)
	fmt.Println("not reverse : ", str32Bit) //  00000000000000000000000000000001
	str32Rune := []rune(str32Bit)
	for i, val := range str32Rune {
		if val == '0' {
			str32Rune[i] = '1'
		} else {
			str32Rune[i] = '0'
		}
	}
	fmt.Println("reversed : ", string(str32Rune))             // 11111111111111111111111111111110
	result, err := strconv.ParseInt(string(str32Rune), 2, 64) // 4294967294
	if err != nil {
		return 0, err
	}
	return result, nil
}
