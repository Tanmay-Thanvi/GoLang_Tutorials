package main

import "fmt"

func main() {
	fmt.Println("Conditional statements")

	loginCount := 23

	var result string

	if loginCount < 10 { // this curly braces must be in that line only 
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("Even no.")
	} else {
		fmt.Println("Odd no.")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num greater or equal than 10")
	}

	// if err != nil {
		 
	// }
}
