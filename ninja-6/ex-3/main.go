package main

import "fmt"

func main() {
	defer foo()

	fmt.Println("Should see me first")
}

func foo() {
	fmt.Println("I'm deferred!")
}
