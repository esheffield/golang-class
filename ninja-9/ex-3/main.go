package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const numRoutines int = 100
	counter := 0

	wg.Add(numRoutines)

	fmt.Println("Before:", counter)

	for i := 0; i < numRoutines; i++ {
		go func() {
			x := counter
			runtime.Gosched()
			x++
			counter = x
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("After: ", counter)
}
