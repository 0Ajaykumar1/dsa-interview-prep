// Given a string str1, Return the array containing the frequency of each character in the order of their occurrence.

// Input Format
// First Parameter - string str1

// Output Format
// Return the array.

package main

import (
	"fmt"
	"strings"
)

func solve(str string) []int {
	freq := make(map[string]int)

	for _, i := range strings.Split(str, "") {
		if _, ok := freq[i]; ok {
			freq[i] += 1
		} else {
			freq[i] = 1
		}
	}

	res := make([]int, len(freq), len(freq))
	i := 0
	for _, val := range freq {
		res[i] = val
		i++
	}
	return res
}

func main() {
	var str string
	fmt.Print("Enter string: ")
	fmt.Scanf("%s", &str)
	res := solve(str)
	fmt.Println(res)
}
