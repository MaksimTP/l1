package main

import (
	"fmt"
	"math"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// This function takes three parameters: (value) that needs to be changed, (pos) - bit that needs to be changed, (toOne) - value of a bit (true - 1, false - 0)
func setBit(value int64, pos int, toOne bool) int64 {
	if pos > 63 || pos < 0 {
		return value
	}
	pos = int(math.Pow(2, float64(pos)))
	if toOne {
		value |= int64(pos)
	} else {
		value |= int64(pos)
		value ^= int64(pos)
	}
	return value
}

func main() {
	var a int64 = 254
	fmt.Printf("a in binary: %b, a in binary after setting bit: %b\n", a, setBit(a, 0, true))
	fmt.Println(a, setBit(a, 0, true))
}
