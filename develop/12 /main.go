package main

import "fmt"

// Implementation of "Set"
type Set struct {
	Arr []string
}

// This function creates a pointer to a Set that forms from slice given as a parametr
func NewSet(arr []string) (*Set, error) {

	uniqueArr := []string{}

	for _, v := range arr {
		isIn := false
		for _, v2 := range uniqueArr {
			if v == v2 {
				isIn = true
			}
		}
		if !isIn {
			uniqueArr = append(uniqueArr, v)
		}
	}

	set := &Set{Arr: uniqueArr}
	return set, nil
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	s, _ := NewSet(arr)
	fmt.Println(s.Arr)
}
