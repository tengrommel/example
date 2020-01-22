package main

import "fmt"

type a struct {
	XX int
	YY int
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

type b struct {
	AA string
	XX int
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

type c struct {
	A a
	B b
}

func main() {
	var i c
	i.A.A()
	i.B.A()
}
