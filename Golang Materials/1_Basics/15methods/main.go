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

	// fn 
	tanmay.GetStatus()
	tanmay.NewEmail()

	fmt.Println(tanmay) // no change bcoz values are passed not reference
	// will do it soon 
}

type User struct { // keep first letter captial (not must but imp)
	Name   string
	Email  string
	Age    int
	Status bool
	// oneAge int // this oneage cant be used as first letter is small
}

func (u User) GetStatus(){
	fmt.Println("Is user Active: ",u.Status)
}
// copied from structs file 

func (u User) NewEmail()  {
	u.Email = "abc@g.com"
	fmt.Println("Mail is ",u.Email)
}