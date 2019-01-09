package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(sh shape) {
	fmt.Println("The area of the shape is:", sh.area())
}

func main() {
	s := square{
		side: 10,
	}

	c := circle{
		radius: 100,
	}

	info(s)
	info(c)
}
