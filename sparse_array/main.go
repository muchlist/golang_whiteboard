package main

import "fmt"

func main() {
	input := []string{"abc", "ab", "ccd", "ddd", "abc", "ccd"}
	queries := []string{"abc", "ab", "ccd"}

	result := matchingStrings(input, queries)
	fmt.Printf("%v", result) // [2 1 2]
}

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 *
 * input : aba baba aba xzxb
 * query : aba xzxb ab
 * result : 2 1 0
 */
func matchingStrings(strings []string, queries []string) []int {
	// membuat slice int berdasarkan query
	result := make([]int, len(queries))
	for i, query := range queries {
		var count int
		// loop strings
		for _, text := range strings {
			if text == query {
				count++
			}
		}
		result[i] = count
	}
	return result
}
