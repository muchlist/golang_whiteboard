package main

import (
	"fmt"
	"strconv"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzbuzz = "fizzbuzz"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("angka : %d >>>> %s\n", i, fizzBuzz(i))
	}
}

// fizzBuzz mengembalikan fizz jika angkanya adalah kelipatan 3 dan buzz jika angka kelipatan 5
// dan fizzbuzz jika angka adalah kelipatan 3 sekaligus 5
func fizzBuzz(input int) string {
	if input < 1 {
		return strconv.Itoa(input)
	}
	if input%3 == 0 && input%5 == 0 {
		return fizzbuzz
	} else if input%3 == 0 {
		return fizz
	} else if input%5 == 0 {
		return buzz
	} else {
		return strconv.Itoa(input)
	}
}
