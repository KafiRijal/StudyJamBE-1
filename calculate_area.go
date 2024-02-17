package main

import (
	"math"
)

/*
	Hint: Use the following formulas to calculate the area of different shapes.
	Rectangle: length * width
	Circle: π * radius * radius
	Triangle: 0.5 * base * height
	Square: side * side

	Note: You can use math.Pi for π and math.Pow(x, y) for x^y.
*/

// Area calculates the area of different shapes based on the shape and provided dimensions.
func Area(shape string, dimensions []float64) float64 {
	// TODO 1: Create a variable named area of type float64 and assign 0 to it.
	var area float64 = 0

	// TODO 2: Use a switch statement or if statement to calculate the area of the shape.
	switch shape {
	case "Rectangle":
		area = dimensions[0] * dimensions[1]
	case "Circle":
		area = math.Pi * math.Pow(dimensions[0], 2)
	case "Triangle":
		area = 0.5 * dimensions[0] * dimensions[1]
	case "Square":
		area = math.Pow(dimensions[0], 2)
	}

	// TODO 3: Return the area variable.
	return area
}
