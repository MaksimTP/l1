package main

import "fmt"

func partition(arr []int, low int, high int) int {

	// Choose the pivot
	pivot := arr[high]

	// Index of smaller element and indicates
	// the right position of pivot found so far
	i := low - 1

	// Traverse arr[low..high] and move all smaller
	// elements on left side. Elements from low to
	// i are smaller after every iteration
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Move pivot after smaller elements and
	// return its position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low int, high int) {
	if low < high {

		// pi is the partition return index of pivot
		pi := partition(arr, low, high)

		// Recursion calls for smaller elements
		// and greater or equals elements
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {

	arr := []int{12, 3, 1, 23, 5, 6, 7, 12}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
