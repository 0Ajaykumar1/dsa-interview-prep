package main

import (
	"fmt"
	"math"
)

// Let us say your expense for every month are listed below,
// January - 2200
// February - 2350
// March - 2600
// April - 2130
// May - 2190
// Create a list to store these monthly expenses and using that find out,

// 1. In Feb, how many dollars you spent extra compare to January?
// 2. Find out your total expense in first quarter (first three months) of the year.
// 3. Find out if you spent exactly 2000 dollars in any month
// 4. June month just finished and your expense is 1980 dollar. Add this item to our monthly expense list
// 5. You returned an item that you bought in a month of April and
// got a refund of 200$. Make a correction to your monthly expense list
// based on this

func main() {
	expenses := [5]int{2200, 2350, 2600, 2130, 2190}

	// 1. In Feb, how many dollars you spent extra compare to January?
	fmt.Printf("Excess spent in February is %v", math.Abs(float64(expenses[0]-expenses[1])))

	// 2. Find out your total expense in first quarter (first three months) of the year.
	fmt.Println("Total first quarter expenses are ", expenses[0]+expenses[1]+expenses[2])

	// 3. Find out if you spent exactly 2000 dollars in any month
	var test = false

	for _, v := range expenses {
		if v == 2000 {
			test = true
		}
	}

	fmt.Println("Is any month expenses are $2000 ", test)

	// 4. June month just finished and your expense is 1980 dollar. Add this item to our monthly expense list

	var newExpenses [len(expenses) + 1]int

	for i, _ := range expenses {
		newExpenses[i] = expenses[i]
	}

	newExpenses[len(expenses)] = 1980

	fmt.Println(newExpenses)

	// 5. You returned an item that you bought in a month of April and
	// got a refund of 200$. Make a correction to your monthly expense list
	// based on this
	newExpenses[3] = newExpenses[3] - 200
	fmt.Println(newExpenses)

}
