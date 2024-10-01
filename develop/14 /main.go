package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

// This function can define type of a variable in a runtime. Can define only int,string,bool and chan; for every other type returns empty string and a error
func getType(a interface{}) (string, error) {

	// reflect.ValueOf(a).Kind().String() returns a string representation of a type of a variable
	if reflect.ValueOf(a).Kind().String() == "int" {
		return "int", nil
	} else if reflect.ValueOf(a).Kind().String() == "string" {
		return "string", nil
	} else if reflect.ValueOf(a).Kind().String() == "bool" {
		return "bool", nil
	} else if reflect.ValueOf(a).Kind().String() == "chan" {
		return "chan", nil
	} else {
		return "", fmt.Errorf("unknown type")
	}
}

func main() {
	x := make(chan int, 2)
	a := []interface{}{2, "asd", true, [5]int{1, 2, 3, 4, 5}, x}
	for _, val := range a {
		t, err := getType(val)
		if err != nil {
			fmt.Println("error:", err.Error(), val)
		} else {
			fmt.Println(val, "->", t)
		}
	}
}
