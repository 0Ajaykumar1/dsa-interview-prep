// You are given two integers, a and b.
// Compute the Sum and Difference of a and b. Return the Product of Sum and Difference.

// Input Format
// First Parameter - Integer a

// Second Parameter - Integer b

// Output Format
// Return the integer.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(a, b int) int {
	// CODE HERE
	return ((a + b) * (a - b))
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number:")
	f, _ := reader.ReadString('\n')
	fmt.Println("Enter second number:")
	s, _ := reader.ReadString('\n')
	f = strings.TrimSpace(f)
	s = strings.TrimSpace(s)
	a, _ := strconv.Atoi(f)
	b, _ := strconv.Atoi(s)

	fmt.Println(solve(a, b))
}
