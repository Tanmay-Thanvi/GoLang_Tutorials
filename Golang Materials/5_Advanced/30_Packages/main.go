package main

import (
	"Learn_Modules/geometry" // Importing the geometry package
	"fmt"
)

func main() {
	// Creating a Rectangle instance
	rect := geometry.Rectangle{Width: 15, Height: 20}
	circle := geometry.Circle{Radius: 2.0}

	geometry.Pi = 3
	geometry.PrintPi()
	// Calculating the area using the package function
	rect_area := rect.Area()
	circle_area := circle.Area()

	// Displaying the calculated area
	fmt.Println("Area of the rectangle:", rect_area)
	fmt.Println("Area of the Circle:", circle_area)
}
