package main

import "fmt"

func main() {
	intSlice := []int32{1, 1, 3, 4, 4, 5, 5}
	result := lonelyInteger(intSlice)

	fmt.Println(result) // expect 3
}

/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */
func lonelyInteger(a []int32) int32 {
	intMap := make(map[int32]int)
	for _, v := range a {
		_, ok := intMap[v]
		if !ok {
			intMap[v] = 1
		} else {
			intMap[v] = intMap[v] + 1
		}
	}
	for key, val := range intMap {
		if val == 1 {
			return key
		}
	}
	return 0
}
