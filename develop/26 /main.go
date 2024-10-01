package main

import "fmt"

func isUnique(s string) bool {
	// Initialization of a hashmap
	seen := make(map[rune]bool)

	for _, v := range s {

		// If there is no value in map, add it to map
		if value, found := seen[v]; !found {
			seen[v] = value
		} else {
			// if there was value in map, then return false, because symbol is not unique
			return false
		}
	}
	return true
}

func main() {
	a1 := "abcd"
	a2 := "aAs"
	a3 := "aas"
	fmt.Println(isUnique(a1))
	fmt.Println(isUnique(a2))
	fmt.Println(isUnique(a3))
}
