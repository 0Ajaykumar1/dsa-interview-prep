package main

import "fmt"

func main() {

	var stocks = [5]int{298, 305, 320, 301, 292}
	var length = len(stocks)
	//lookup/searching an element in array
	for i := 0; i < length; i++ {
		if stocks[i] == 320 {

			fmt.Println("Searching an element in an array \n", i)
		}
	}

	// traversal of array
	fmt.Println("Traversal of an array")
	for _, stockPrice := range stocks {
		fmt.Println(stockPrice)
	}

	//adding an element to array
	fmt.Println("Adding an element to array")
	var newStocks [len(stocks) + 1]int
	for i := 0; i < length; i++ {
		newStocks[i] = stocks[i]
	}
	newStocks[length] = 840

	fmt.Println(stocks)
	fmt.Println(newStocks)

	// inserting an element into array
	fmt.Println("Inserting an element into array")
	var insertedStocks [len(stocks) + 1]int
	index := 3
	value := 420
	for i := 0; i < length+1; i++ {
		if i < index {
			insertedStocks[i] = stocks[i]
		} else if i == index {
			insertedStocks[i] = value
		} else {
			insertedStocks[i] = stocks[i-1]
		}
	}

	fmt.Println(insertedStocks)

	// deleting an element from array
	fmt.Println("Deleting an element from an array")
	var delStocks [len(stocks) - 1]int
	for i := 0; i < length-1; i++ {
		if i >= index {
			delStocks[i] = stocks[i+1]
		} else {
			delStocks[i] = stocks[i]
		}
	}

	fmt.Println(delStocks)

}
