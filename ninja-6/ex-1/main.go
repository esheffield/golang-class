package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println("foo: ", x)
	fmt.Println("bar:", y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 99, "bottles of beer on the wall"
}
