package main

import "fmt"

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
	t1 := triangle{
		height: 10.0, base: 10.0,
	}

	s1 := square{
		sideLength: 2.0,
	}

	t1.getArea()
	s1.getArea()

	printArea(t1)
	printArea(s1)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
