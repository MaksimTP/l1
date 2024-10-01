package main

import "fmt"

// Implementation of a removal function. Removes a value by index. Index is [0, len(arr) -1]
func remove(arr []int, ind int) ([]int, error) {
	if ind < 0 || ind >= len(arr) {
		return arr, fmt.Errorf("index out of range")
	}
	return append(arr[:ind], arr[ind+1:]...), nil
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	sl, err := remove(sl, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sl)
	}
}
