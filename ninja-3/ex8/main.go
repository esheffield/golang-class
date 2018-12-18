package main

import "fmt"

func main() {
	x := "Dr No"
	switch {
	case x == "James Bond":
		fmt.Println("Good evening, Mr. Bond")
	case x == "Moneypenny":
		fmt.Println("Hello Ms. Moneypenny")
	default:
		fmt.Println("Are you a spy?")
	}
}
