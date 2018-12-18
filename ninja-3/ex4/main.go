package main

import "fmt"

func main() {
	y := 1967
	for {
		fmt.Println(y)
		y++
		if y > 2018 {
			break
		}
	}
}
