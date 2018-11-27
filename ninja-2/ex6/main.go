package main

import "fmt"

const (
	y2015 = 2015 + iota
	y2016
	y2017
	y2018
)

func main() {
	fmt.Println(y2015, y2016, y2017, y2018)
}
