package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	const numRoutines int = 100
	counter := 0

	wg.Add(numRoutines)

	fmt.Println("Before:", counter)

	for i := 0; i < numRoutines; i++ {
		go func() {
			mu.Lock()
			x := counter
			x++
			counter = x
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("After: ", counter)
}
