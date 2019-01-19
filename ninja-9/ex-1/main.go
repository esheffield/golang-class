package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Starting")
	go hiOne()
	go hiTwo()
	fmt.Println("Waiting")
	wg.Wait()
	fmt.Println("All done!")
}

func hiOne() {
	fmt.Println("Hello 1")
	wg.Done()
}

func hiTwo() {
	fmt.Println("Hello 2")
	wg.Done()
}
