package main

import "fmt"

func groupValues(arr []float32, step int) map[int][]float32 {
	m := make(map[int][]float32)

	for _, v := range arr {
		// Get a key for a value by rounding to the lower ten
		key := int(v) / step * step

		// Append value to a slice with specified key
		m[key] = append(m[key], v)
	}
	return m
}

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupValues(arr, 10))
}
