package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}

	fooSum := foo(xi...)
	barSum := bar(xi)

	fmt.Println(fooSum)
	fmt.Println(barSum)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
