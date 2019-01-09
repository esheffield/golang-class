package main

import "fmt"

func main() {
	addFive := adder(5)

	fmt.Println("Result: ", addFive(7))
}

func adder(i int) func(int) int {
	return func(x int) int {
		return i + x
	}
}
