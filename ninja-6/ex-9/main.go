package main

import "fmt"

func main() {
	f := getGreeter()

	runGreeter(f)
}

func runGreeter(greeter func()) {
	greeter()
}

func getGreeter() func() {
	return func() {
		fmt.Println("Hello from returned callback func")
	}
}
