package main

import "fmt"

type shape interface {
	getArea() string
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{
		sideLength: 4,
	}
	tr := triangle{
		height: 4,
		base:   5,
	}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() string {
	area := fmt.Sprintf("Square area: %f", s.sideLength*s.sideLength)
	return area
}

func (t triangle) getArea() string {
	area := fmt.Sprintf("Triangle area: %f", 0.5*t.base*t.height)
	return area
}
