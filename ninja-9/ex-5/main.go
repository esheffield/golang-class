package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	const numRoutines int = 100
	var counter int64

	wg.Add(numRoutines)

	fmt.Println("Before:", counter)

	for i := 0; i < numRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("After: ", counter)
}
