package main

import "fmt"

// This function implement reversion of a string. Works with runes.
func reverse(s string) string {
	r := []rune(s)
	res := ""
	for _, v := range r {
		res = string(v) + res
	}
	return res
}

func main() {
	x := "hell😊o"
	fmt.Println(x, reverse(x))
}
