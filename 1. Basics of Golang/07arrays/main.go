package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays lec")

	var fruits [4]string
	fruits[0] = "aPPLE"
	fruits[1] = "Banana"
	fruits[3] = "Peach" 

	fmt.Println("List : ",fruits)
	fmt.Println("Length : ",len(fruits))

	var veglist = [3]string{"Potato","Beans","Mushroom"}
	fmt.Println("List : ",veglist)
	fmt.Println("Len : ",len(veglist))
}
