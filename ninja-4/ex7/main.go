package main

import "fmt"

func main() {
	a := [][]string{
		{"James", "Bond", "Shaken, not stirred."},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, v := range a {
		for _, n := range v {
			fmt.Printf("%s\t", n)
		}
		fmt.Println("")
	}
	fmt.Println(a)
}
