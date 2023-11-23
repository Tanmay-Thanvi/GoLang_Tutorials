package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello"
	fmt.Println(welcome)

	// pkg.go.dev 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza : ")

	// comma ok syntax || err ok 
	input, _ := reader.ReadString('\n') // we can write input , error (it is just like try and catch ) 
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("Type of rating is %T",input)
	
}
