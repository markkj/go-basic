package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	width  float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{20, 15}
	s := square{20}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.height * t.width * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
