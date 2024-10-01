package main

import (
	"fmt"
	"main/distance"
)

func main() {
	p1 := distance.MakePoint(1, 1)
	p2 := distance.MakePoint(0, 0)
	line := distance.Line{P1: p1, P2: p2}
	fmt.Println(line.Length())
}
