package main

import "fmt"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%d\n", int64(i))
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
