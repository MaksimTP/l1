package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := "snow dog sun"

	// Get slice from string separated by " "
	arr := strings.Split(s, " ")
	// Reverse slice
	slices.Reverse(arr)
	// Get string from slice by joining its elements with sep=" "
	res := strings.Join(arr, " ")
	fmt.Println("Original:", s, "->", "Reversed:", res)
}
