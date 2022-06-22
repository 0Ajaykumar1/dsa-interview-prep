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
	var rArr []int
	if y == len(arr)-1 {
		rArr = arr[x:]
	} else {
		rArr = arr[x : y+1]
	}
	for i := 0; i < len(rArr); i++ {
		sum += rArr[i]
	}
	return (float64(sum) / float64(len(rArr)))
}

func main() {
	res := solve(6, []int{6, 2, 5, 4, 3, 1}, 2, 5)
	fmt.Println(res)
}
