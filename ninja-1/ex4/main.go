package main

import "fmt"

type magic int

var x magic

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
