package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

/*
struct Human:

	atributes:
		- name -- string, holds a name of a Human
		- age -- int, holds an age of a Human
	methods:
		- (*Human) changeName(string)
		- (*Human) changeAge(int)
*/
type Human struct {
	name string
	age  int
}

// Action is a structure that embed Human and has additional field job.
type Action struct {
	Human
	job string
}

// change name of a human object
func (h *Human) changeName(name string) {
	h.name = name
}

// change age of a human object. age can only be changed if it is >= 0.
func (h *Human) changeAge(age int) {
	if age >= 0 {
		h.age = age
	}
}

func main() {
	worker := &Action{
		Human: Human{name: "Tom", age: 20},
		job:   "Actor",
	}
	worker.changeAge(30)
	fmt.Println(worker.age, worker.job)
}
