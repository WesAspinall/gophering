package main

import (
	"fmt"
	"math"
)

//method sets are the methods attached to a type
type square struct {
	width float64
}

type circle struct {
	radius float64
}

// see where in info() is called in func main

// non-pointer receiver, works with pointer or non-pointer values
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// pointer receiver, only works with values that are pointers
func (s *square) area() float64 {
	return s.width * s.width
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s1 := square{
		width: 4,
	}

	c1 := circle{2} // clean af

	fmt.Println(info(&s1))
	fmt.Println(info(c1))
}
