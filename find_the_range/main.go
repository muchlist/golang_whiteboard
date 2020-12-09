package main

import (
	"fmt"
	"github.com/muchlist/golang_whiteboard/utils"
)

func main() {

	/*
		Write a function that searches a sorted ArrayList and finds the range of particular element.
		Input :
		    [1,2,3,3,3,4,5,5,5], 3
		Output :
		    2 to 4
	*/
	data := []int{1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 6}
	valueToSearch := 3

	// cara pertama (kotlin whiteboard)
	// ada bug pada findIndices
	startIndex, endIndex := findIndices(data, valueToSearch)
	fmt.Printf("%d to %d\n", startIndex, endIndex)

	// cara kedua
	// cari index yang diketahui menggunakan binary search
	indexKnown, err := utils.BinarySearch(data, valueToSearch)
	if err != nil {
		fmt.Println("index not found")
	} else {
		// berdasarkan index yang diketahui, cari index awal dan akhir
		startIndex := firstPositionIterate(data, indexKnown)
		endIndex := lastPositionIterate(data, indexKnown)
		fmt.Printf("%d to %d\n", startIndex, endIndex)
	}
}

// --------------------------Cara Pertama-------------------------------

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

// --------------------------Cara Pertama End-------------------------------

// --------------------------Cara Kedua-------------------------------

// mencari index awal dengan perulangan menggunkan index yang sudah diketahui
// kelemahannya adalah apabila nilai dengan index yang sama ada banyak
func firstPositionIterate(sortedList []int, index int) int {
	firstValue := sortedList[index]
	firstIndex := index
	for firstIndex != 0 && firstValue == sortedList[firstIndex-1] {
		firstIndex--
	}
	return firstIndex
}

// mencari index akhir dengan perulangan menggunkan index yang sudah diketahui
// kelemahannya adalah apabila nilai dengan index yang sama ada banyak
func lastPositionIterate(sortedList []int, index int) int {
	firstValue := sortedList[index]
	firstIndex := index
	maxIndex := len(sortedList) - 1

	for firstIndex != maxIndex && firstValue == sortedList[firstIndex+1] {
		firstIndex++
	}
	return firstIndex
}

// --------------------------Cara Kedua End-------------------------------
