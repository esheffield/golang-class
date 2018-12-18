package main

import "fmt"

func main() {
	if x := 1; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
}
