package main

import {
	"math"
}

// create an interface shape
type shape interface {
	area() float64
	//pow(x int) float64
}

// create a type Circle
type Circle struct {
	radius float64
}

// create a type Rectangle
type Rect struct {
	width float64
	height float64
}

// function to calcualte the area for circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// function to calcualte the area for rectangle
func (r Rect) struct {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	r1 := Rect{5, 7}
	shapes := []shape{c1, r1} // because both Circle and Rect have the area() method
	                          // so c1 & r1 can belong to the shape interface
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
	// Output; 
	// 63.617
	// 35

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
	// Output; 
	// 63.617
	// 35
}
