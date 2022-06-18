package main

import (
	"fmt"
)

/*
Temukan dua integer yang berbeda, yang apabila dijumlahkan akan menghailkan
angka sesuai target


input :
array = [10,5,2,3,-6,9,11]
s = 4

output :
[10, -6]
*/

func main() {
	input := []int{10, 5, 2, 3, -6, 9, 11}
	s := 4

	result := pairSum(input, s)
	fmt.Printf("input\t: %v\ntarget\t: %d\nresult\t: %v\n", input, s, result)

	resultOptimized := pairSumOptimized(input, s)
	fmt.Printf("input\t: %v\ntarget\t: %d\nresult\t: %v\n", input, s, resultOptimized)
}

// N^2
func pairSum(ar []int, sum int) []int {

	// lakukan perulangan untuk setiap nilai pada ar
	for i := 0; i < len(ar); i++ {

		// nilai pertama
		val := ar[i]

		// lakukan perulangan untuk setiap nilai pada ar
		// yang berbeda nilai
		for i := 0; i < len(ar); i++ {

			// nilai kedua
			val2 := ar[i]

			// jika nilainya sama continue
			if val == ar[i] {
				continue
			}

			// Jumlahkan val1 dan val2 apabila hasilnya sama dengan target,
			// return
			if val+val2 == sum {
				return []int{val, val2}
			}
		}
	}

	return nil
}

func pairSumOptimized(ar []int, sum int) []int {
	mapping := make(map[int]int)

	for i := 0; i < len(ar); i++ {

		// x adalah pasangan angkanya
		x := sum - ar[i]
		// cari pasangan angka tersebut di map
		if c, ok := mapping[x]; ok {
			return []int{ar[c], ar[i]}
		}
		// jika tidak menemukan maka simpan index di map dengan key angka
		mapping[ar[i]] = i
	}
	return nil
}
