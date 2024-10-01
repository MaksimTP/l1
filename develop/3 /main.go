package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	// Initialize waitgroup object
	var wg sync.WaitGroup

	var res *int = new(int)

	for i := 0; i < len(arr); i++ {
		// add 1 to the counter of wg object
		wg.Add(1)
		go func(val int, wg *sync.WaitGroup) {
			// Decrements a counter by one of a wg object
			defer wg.Done()
			// Dereference pointer res and add a square of val
			*res += val * val
		}(arr[i], &wg)
	}

	// blocks a goroutine until wg counter is zero
	wg.Wait()
	fmt.Println(*res)

}
