// You are given a String str. Return the String by Concatenating it with String "Hello, " and an exclamation mark at the end “!”.

// Input Format
// First Parameter - String str

// Output Format
// Return a String

package main

import (
	"fmt"
)

func solve(str string) string {
	res := "Hello " + str + "!"
	return res
}

func main() {
	var str string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &str)
	res := solve(str)
	fmt.Println(res)
}
