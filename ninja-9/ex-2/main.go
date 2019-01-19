package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func (p *person) speak() {
	fmt.Println("My name is", p.First, p.Last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	fmt.Println("The human says:")
	h.speak()
}

func main() {
	p := person{
		First: "Fred",
		Last:  "Flintstone",
		Age:   34,
	}

	// This doesn't work
	// saySomething(p)
	saySomething(&p)
}
