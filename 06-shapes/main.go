package main

import "fmt"

// SQUARE
type square struct {
	side float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

// TRIANGLE
type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

// SHAPE
type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

// MAIN
func main() {
	s := square{2}
	t := triangle{1, 3}
	printArea(s)
	printArea(t)
}
