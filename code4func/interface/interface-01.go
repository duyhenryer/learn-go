package main

import "fmt"

// Define an interface named "Shape"
type Shape interface {
	Area() float64
}

// Implement the Shape interface for a Rectangle  (Hình chữ nhật )
type Rectangle struct {
	With   float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.With * r.Height
}

// Implement the Shape interface for a Circle type
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function that takes a Shape interface and prints its area
func printArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func main() {
	// Create instance of Rectangle and Circle
	rectangle := Rectangle{With: 11, Height: 18}
	circle := Circle{Radius: 2}
	// Call the printArea function with instances of different types
	printArea(rectangle)
	printArea(circle)
}
