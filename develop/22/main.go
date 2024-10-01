package main

import (
	"fmt"
	"math/big"
)

// This function allows to work with big.Ints and do some operations with them.
// Allowed operations: "add", "plus", "sub", "minus", "mul", "div".
// Returns calculated value with a type big.Int
func Operation(a *big.Int, b *big.Int, oper string) big.Int {
	switch oper {
	case "add", "plus":
		return *new(big.Int).Add(a, b)
	case "sub", "minus":
		return *new(big.Int).Add(a, b)
	case "mul":
		return *new(big.Int).Mul(a, b)
	case "div":
		return *new(big.Int).Div(a, b)
	default:
		return *big.NewInt(0)
	}

}

func main() {
	// Create to big.Ints
	a := big.NewInt(4294967296) // 2 ^ 32
	b := big.NewInt(2147483648) // 2 ^ 31
	res := Operation(a, b, "mul")
	fmt.Println(res)

}
