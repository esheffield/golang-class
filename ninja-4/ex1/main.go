package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
