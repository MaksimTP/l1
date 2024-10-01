package main

import "fmt"

// This function swaps values of two variables
func swapValues(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {

	a := 2
	b := 3
	fmt.Println(a, b)
	swapValues(&a, &b)
	fmt.Println(a, b)
}
