package main

import "fmt"

func main() {
	x := "Dr No"
	if x == "James Bond" {
		fmt.Println("Good evening, Mr. Bond")
	} else if x == "Moneypenny" {
		fmt.Println("Hello Ms. Moneypenny")
	} else {
		fmt.Println("Are you a spy?")
	}
}
