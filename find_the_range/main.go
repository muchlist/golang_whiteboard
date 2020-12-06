package main

import (
	"fmt"
	"github.com/muchlist/golang_whiteboard/utils"
)

func main() {
	data := []int{1, 5, 30, 70, 80, 100, 120}

	index, err := utils.BinarySearchRecursive(data, 120, 0, len(data)-1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	index2, err := utils.BinarySearch(data, 120)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("index 1: %d\nindex 2: %d", index, index2)
}
