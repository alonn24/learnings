package main

import "log"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Structures - Structures
func structures() {
	var r Shape = Rectangle{Length: 3, Width: 4}
	log.Printf("> structures: Type of r: %T, Area: %v, Perimeter: %v.\n", r, r.Area(), r.Perimeter())
}
