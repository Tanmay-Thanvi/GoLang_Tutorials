package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// No inheritance in golang; No parent / super
	tanmay := User{"Tanmay","tanny@gmail.com",12,true} // brackets are curly here
	fmt.Println(tanmay)
	// different in both print statements 
	fmt.Printf("Details are %+v\n",tanmay) // +v gives more details 
	fmt.Printf("Name is %v and email is %v\n",tanmay.Name,tanmay.Email)
}

type User struct { // keep first letter captial (not must but imp)
	Name   string
	Email  string
	Age    int
	Status bool
}
