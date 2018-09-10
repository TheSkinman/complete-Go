package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tri := triangle{10, 10}
	sqr := square{10}

	printArea(tri)
	printArea(sqr)
}

func printArea(s shape) {
	fmt.Printf("The area of the %T is %v\n", s, s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
