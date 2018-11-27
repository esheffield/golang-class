package main

import "fmt"

type magic int

var x magic
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
