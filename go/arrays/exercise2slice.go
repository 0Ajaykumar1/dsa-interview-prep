package main

import (
	"fmt"
	"sort"
)

func main() {
	heros := []string{"spider man", "thor", "hulk", "iron man", "captain america"}

	// 1. Length of the list;
	fmt.Println("length of the heros list is ", len(heros))

	// 2. Add 'black panther' at the end of this list
	heros = append(heros, "black panther")
	fmt.Println("list after adding black panther ", heros)

	// 3. You realize that you need to add 'black panther' after 'hulk',
	// so remove it from the list first and then add it after 'hulk'

	heros = heros[:5]
	heros = append(heros[:4], heros[3:]...)
	heros[3] = "black panter"
	fmt.Println("list after add black panther after hulk", heros)

	// 4. Now you don't like thor and hulk because they get angry easily :)
	// So you want to remove thor and hulk from list and replace them with doctor
	// strange (because he is cool).
	// Do that with one line of code.

	heros = append(append(heros[:1], "doctor strange"), heros[3:]...)
	fmt.Println("list after replacing thor and hulk with doctor strange ", heros)

	// 5. Sort the heros list in alphabetical order (Hint. Use dir() functions to
	// list down all functions available in list)

	sort.Slice(heros, func(i, j int) bool {
		return heros[i] < heros[j]
	})
	fmt.Println("list after sorting ", heros)
}
