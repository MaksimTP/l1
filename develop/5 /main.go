package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// create buffered channel
	ch := make(chan int, 5)

	var N int = 1
	// create channel that will emit signal after N seconds
	timeout := time.After(time.Duration(N) * time.Second)

	go func() {
		// cycle that will read values until signal from timeout is received
		for {
			select {
			case x := <-ch:
				fmt.Println(x)
			case <-timeout:
				fmt.Println("timeout happened, exiting...")
				os.Exit(0)
			}
		}
	}()

	// send values to channel
	for {
		ch <- rand.Int()
	}

}
