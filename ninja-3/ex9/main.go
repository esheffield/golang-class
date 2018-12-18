package main

import "fmt"

func main() {
	favSport := "Golf"

	switch favSport {
	case "Curling":
		fmt.Println("Obviously a person of taste!")
	case "Football":
		fmt.Println("Hooligan, obviously (American)")
	case "Soccer":
		fmt.Println("Hooligan, obviously (non-American)")
	default:
		fmt.Println("Neeeeeerd!")
	}
}
