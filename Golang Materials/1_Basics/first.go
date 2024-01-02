package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string = "Tanmay"
	var age int = 21
	var Details = name + " - " + strconv.Itoa(age)
	fmt.Println("My Details :\n", Details)
}
