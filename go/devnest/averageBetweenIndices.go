// Given an array arr of size n, and two intervals x and y. Find the average of elements which lies between the given intervals inclusively.

// Input Format
// First Parameter - number n

// Second Parameter - array arr of size n

// Third Parameter - number x

// Fourth Parameter - number y

// Output Format
// Return the decimal value

package main

import "fmt"

func solve(n int, arr []int, x int, y int) float64 {
	sum := 0
	for i := x; i <= y; i++ {
		sum += arr[i]
	}
	return (float64(sum) / float64(y-x+1))
}

func main() {
	res := solve(6, []int{6, 2, 5, 4, 3, 1}, 2, 5)
	fmt.Println(res)
}
