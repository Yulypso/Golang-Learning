package main

import "fmt"

/* struct definition */
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

/* interface definition */
type shape interface {
	getArea() float64
}

/* main func */
func main() {
	myTriangle := triangle{
		height: 12.0,
		base:   6.0,
	}

	mySquare := square{}
	mySquare.sideLength = 5.0

	printArea(myTriangle)
	printArea(mySquare)
}

/* Thanks to these func, triangle and square are now shape type */
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

/* shape func definition */
func printArea(s shape) {
	fmt.Println("Shape Area:", s.getArea())
}
