package main

import (
	"fmt"
)

var justString string

// This function creates big string
func createHugeString(num int) (res string) {

	for i := 0; i < num; i++ {
		res += "h"
	}
	return
}

// In this function we say that var justString is equal to a slice of a string v. That creates some leaky problems. We still have allocated space for string with len (1 << 13) but we have access to the first 100 bytes. So, while justString is alive, we have a leak about (1 << 13) - 100 bytes.

// Solution without leaks would look like: justString = bytes.Clone(v[:100])

func someFunc() {
	v := createHugeString(1 << 13)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}
