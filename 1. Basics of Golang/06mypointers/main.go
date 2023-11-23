package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")
	
	var ptr *int 

	fmt.Println("Value of pointer is : ",ptr) // it will be nil 

	myNum := 23

	ptr = &myNum // or var ptr = &myNum 
	fmt.Println("Value of pointer is : ",ptr)
	fmt.Println("Value stored in pointer is : ",*ptr)

	*ptr *= 2
	fmt.Println("New value is ",*ptr)
}
