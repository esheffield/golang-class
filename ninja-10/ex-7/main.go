package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	q := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- (i * 10) + j
			}
			q <- 1
		}(i)
	}

	receive(c, q, 10)

	fmt.Println("exiting")
}

func receive(c, q <-chan int, quit int) {
	cnt := 0
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case m := <-q:
			cnt += m
			if cnt >= quit {
				return
			}
		}
	}
}
