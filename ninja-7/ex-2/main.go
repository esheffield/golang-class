package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(pony *person) {
	pony.first = "Pinkie"
	pony.last = "Pie"
}

func main() {
	pony := person{
		"Twilight",
		"Sparkle",
	}

	fmt.Println("Before: ", pony)
	changeMe(&pony)
	fmt.Println("After: ", pony)
}
