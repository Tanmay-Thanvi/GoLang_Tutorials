package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	// Adding the values
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages : ", languages)
	fmt.Println("RB Shorts for : ",languages["RB	"])

	delete(languages,"RB")
	fmt.Println(languages)

	// loops in go for maps 

	for _,value := range languages{ // comma ok syntax 
		fmt.Println(" - ",value)
	}
	
}
