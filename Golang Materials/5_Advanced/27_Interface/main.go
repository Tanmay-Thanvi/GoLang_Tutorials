package main

import (
	"fmt"
	"math"
	"strings"
)

// Shape Interface
type shape interface {
	area() float64
	circumference() float64
}

// circle and square structs
type square struct {
	side float64
}

type circle struct {
	radius float64
}

// Square functions
func (s square) area() float64 {
	return s.side * s.side
}

func (s square) circumference() float64 {
	return 4 * s.side
}

// Methods associated with circle (not functions)
func (c circle) area() float64 {
	return c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

// Common function
func printInfo(s shape) {
	fmt.Println("Area : ", s.area())
	fmt.Println("Circumference : ", s.circumference())
}

func printSeperator(s string)  {
	separator := strings.Repeat(s, 20)
	fmt.Println(separator)
}

func main() {

	shapes := []shape{
		circle{radius: 1.2},
		square{side: 5.6},
	}

	for k, v := range shapes {
		printSeperator("-")
		fmt.Printf("K = %v, V = %v \n", k, v)
		printInfo(v)
	}

	printSeperator("-")
}
