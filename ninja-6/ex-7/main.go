package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is an anon func called by variable")
	}

	f()
}
