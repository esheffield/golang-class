package main

import "fmt"

func main() {
	f := getGreeter()

	f()
}

func getGreeter() func() {
	return func() {
		fmt.Println("Hello from returned func")
	}
}
