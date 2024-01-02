package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case i n golang")

	// rand.Seed(time.now().UnixNano()) this is giving error
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ",dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Value is 1 and u can open")	
	case 2:
		fmt.Println("Value is 2 and u can move 2 spots")	
	case 3:
		fmt.Println("Value is 3 and u can move 3 spots")
		fallthrough // If we want fallthrough it will print 3 as well as 4 	
	case 4:
		fmt.Println("Value is 4 and u can move 4 spots")	
	case 5:
		fmt.Println("Value is 5 and u can move 5 spots")	
	case 6:
		fmt.Println("Value is 6 and u can move 6 spots")	
	default:
		fmt.Println("What was that?")
	}
}
