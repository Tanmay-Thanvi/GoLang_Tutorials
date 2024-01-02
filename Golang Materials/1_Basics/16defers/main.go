package main

import "fmt"

func main() {
	defer fmt.Println("World") // it will always execute at last of function 
	defer fmt.Println("One") // LIFO (Last in First Out)
	defer fmt.Println("Two") 

	fmt.Println("Hello")

	mydefer()
}

func mydefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
