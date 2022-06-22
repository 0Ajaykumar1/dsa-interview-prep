// Given a Positive Integer n, return an array of string containing the palindromic string of intergers.

// Input Format
// First Parameter - number n

// Output Format
// Return the array of string.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(n int) []int {
	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, int(math.Pow((math.Pow10(i)-1)/9, 2)))
	}
	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number:")
	f, _ := reader.ReadString('\n')
	f = strings.TrimSpace(f)
	n, _ := strconv.Atoi(f)

	fmt.Println(solve(n))
}
