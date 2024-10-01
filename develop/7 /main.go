package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	// Initialize a mutex object
	var l sync.Mutex

	m := make(map[string]int, 5)

	// Initialize a wg object
	var wg sync.WaitGroup

	// keys for map
	arrKeys := []string{"a", "b", "c", "d", "e"}
	// values for map
	arrValues := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arrKeys); i++ {
		// add 1 to the counter of wg object
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			// block access for other goroutines to avoid overlay writes
			l.Lock()
			// Decrements a counter by one of a wg object
			defer wg.Done()
			// unblock access after the writing to the map is done
			defer l.Unlock()
			m[arrKeys[i]] = arrValues[i]
		}(&wg)
	}
	// blocks a goroutine until wg counter is zero
	wg.Wait()
	fmt.Println(m)
}
