package main

import "fmt"

// Implementation of "Set"
type Set struct {
	Arr []int
}

// This function checks whether value is in set
func (s Set) isIn(value int) bool {
	for _, v := range s.Arr {
		if v == value {
			return true
		}
	}
	return false
}

// This function does intersection of two sets and returns a result of intersection - new set
func (s Set) Intersection(s2 Set) Set {
	newS := Set{}
	for _, v := range s.Arr {
		if !newS.isIn(v) {
			newS.Arr = append(newS.Arr, v)
		}
	}
	for _, v := range s2.Arr {
		if !newS.isIn(v) {
			newS.Arr = append(newS.Arr, v)
		}
	}

	return newS
}

func main() {
	s1 := Set{}
	s1.Arr = []int{1, 2, 3, 4, 5, 6}

	s2 := Set{}
	s2.Arr = []int{2, 3, 4, 5, 6, 7, 3213, 4, 5}

	s3 := s1.Intersection(s2)
	fmt.Println(s3)
}
