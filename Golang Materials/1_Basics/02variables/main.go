package main

import "fmt"

const LoginToken string = "ssbksbxs" // first letter capital means public

func main() {
	var username string = "Tanmay"
	// type fp for shortcut of fmt.Prinrln
	fmt.Println("Username : " + username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println("Status of Logged In : ",isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallVal byte = 255 // byte is an alias for uint8 
	fmt.Println(" Value of smallVal : ",smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.4559445422 // check for float64 as well (Precesion gets extended)
	fmt.Println(" Value of smallFloat : ",smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// Default values and some aliases 
	var anotherVariable int // takes 0 and not any garbage value 
	fmt.Println("Value : ",anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	// implicit type 
	var website = "google.com"
	fmt.Println("Web : "+website)

	// no var style 
	noOfUsers := 30 // it is allowed only inside functions and not at global level
	fmt.Println("No of Users : ",noOfUsers)

	fmt.Println("Login Token : ",LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)

}

// Randomly added for merging purpose