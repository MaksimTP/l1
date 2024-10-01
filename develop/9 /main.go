package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Initializing two (in and out) unbuffered channels
	in := make(chan int)
	out := make(chan int)

	// Lambda function that reads a value from in channel, square it and sends it to out channel
	go func() {
		for {
			res := <-in
			fmt.Printf("source: %d ", res)
			out <- res * res
		}
	}()

	// Lambda function that sends a value to a channel
	go func() {
		for {
			in <- rand.Int() % 15
			time.Sleep(time.Second)
		}
	}()

	// Cycle that reads values from out channel and prints them to a stdout
	for v := range out {
		fmt.Printf("squared: %d\n", v)
	}
	// for {
	// 	fmt.Printf("squared: %d\n", <-out)
	// }

}
