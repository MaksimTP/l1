package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	// Initialize waitgroup object
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		// add 1 to the counter of wg object
		wg.Add(1)

		// Run concurrently a function that prints a double of a given value
		go func(val int, wg *sync.WaitGroup) {
			// Decrements a counter by one of a wg object
			defer wg.Done()
			fmt.Println(val * val)
		}(arr[i], &wg)
	}

	// blocks a goroutine until wg counter is zero
	wg.Wait()
}
