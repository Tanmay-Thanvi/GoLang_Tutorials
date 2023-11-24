package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in Golang")
	// main is called automatically
	greeter()

	// func greeterto()  {
	// 	fmt.Println("Another Method")
	// } // can not do this

	greeterto()

	// we can create anonymous functions (){}
	
	result := adder(3,5)
	fmt.Println("Result is :",result)

	prores, mymsg := proAdder(1,2,4,4,5)
	fmt.Println("Pro Result : ",prores,mymsg)
}

func greeter()  {
	fmt.Println("Namaste from Golang!")
}

func greeterto()  {
		fmt.Println("Another Method")
} 

func adder(val1 int,val2 int) int { // function signature 
	return val1+val2
}

// ... => variadic fns 
func proAdder(values ...int) (int,string) {
	total := 0
	for _,val := range values{
		total += val;
	}
	return total , "Hi Prores"
}