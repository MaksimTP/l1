package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Get number of workers from stdin
	numWorkers := 0
	fmt.Println("Choose number of workers: ")
	fmt.Scan(&numWorkers)

	// Make a chan that will catch sigint signal
	sigs := make(chan os.Signal, 1)

	// catch only sigint signals
	signal.Notify(sigs, syscall.SIGINT)

	ch := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {

		go func() {
			// Receive values from channel until it closed
			for v := range ch {
				select {
				// catch sigint signal
				case <-sigs:
					fmt.Println("SIGINT!")
					close(ch)
					os.Exit(1)
				default:
					break
				}
				// print values to stdout from channel
				fmt.Printf("Num:\t%2d, Worker:\t%2d\n", v, i+1)
			}
		}()
	}
	// send data to channel
	for {
		ch <- rand.Int() % 100
	}

}
