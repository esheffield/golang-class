package main

import (
	"fmt"
)

type customErr struct {
	s string
}

func (e customErr) Error() string {
	return e.s
}

func foo(err error) {
	fmt.Println(err.Error())
}

func main() {
	err := customErr{"I'm a custome error!"}

	foo(err)
}
