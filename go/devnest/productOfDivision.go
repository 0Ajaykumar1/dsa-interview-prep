// You are given two integers, a and b.

// Let x be the integer division and y be the decimal(float) division of a and b.

// Return the product of x and y.

// Input Format
// First Parameter - Integer a

// Second Parameter - Integer b

// Output Format
// Return the decimal product of integer division x and decimal division y of a and b.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(a, b int) float64 {
	intergerDiv := a / b
	floatDiv := float64(a) / float64(b)
	return (float64(intergerDiv) * (floatDiv))
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
