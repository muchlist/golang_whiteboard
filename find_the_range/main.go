package main

import "fmt"

// Mencari posisi pertama index
func firstPosition(sortedList []int, value int) int {
	n := len(sortedList)
	firstPosition := -1

	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		// jika nilai tengah lebih dari atau sama dengan value
		// firstPosition akan di isi dengan index lalu high akan dikurangi 1
		// sampai high akan kurang dari low
		if sortedList[mid] >= value {
			firstPosition = mid
			high = mid - 1
		} else { // sortedList[mid] < value
			low = mid + 1
		}
	}
	return firstPosition
}

// Logika ini akan salah apabila nilai yang dicari ada di akhir slice
// karena fungsi firstPosition hanya mencari index pertama index sedangkan endIndex
// hanya menambahkan 1 angka yang dicari
func findIndices(sortedList []int, value int) (int, int) {
	startIndex := firstPosition(sortedList, value)
	// endIndex menambahkan +1 pada angka yang dicari yang sama dengan angka selanjutnya
	// dengan asumsi angkanya berurutan
	endIndex := firstPosition(sortedList, value+1)

	return startIndex, endIndex - 1
}

func main() {

	/*
		Write a function that searches a sorted ArrayList and finds the range of particular element.
		Input :
		    [1,2,3,3,3,4,5,5,5], 3
		Output :
		    2 to 4
	*/
	data := []int{1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 6}

	startIndex, endIndex := findIndices(data, 3)
	fmt.Printf("%d to %d", startIndex, endIndex)

	// BEST
	// Find the index use binary search
	// Get first position
	// Get last position
}
