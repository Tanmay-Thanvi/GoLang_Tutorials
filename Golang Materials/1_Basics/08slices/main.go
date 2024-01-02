package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("------Slices------")

	var fruitlist = []string{"Apple","Mango"} // in this syntax we always need to initialize 
	fmt.Printf("Type of fruitlist is %T\n",fruitlist)

	fruitlist = append(fruitlist, "Chikoo", "Banana")
	fmt.Println("Feuits : ",fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fruitlist = append(fruitlist[:3])
	fruitlist = append(fruitlist[1:3]) // it will be used alot in development 
	fmt.Println(fruitlist)

	highscore := make([]int,4)
	highscore[0] = 234
	highscore[1] = 989
	highscore[2] = 657
	highscore[3] = 821
	// highscore[4] = 243 this will give error but....

	highscore = append(highscore, 555,65,224)
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	// How Remove a value from slice based on index
	var courses = []string{"Python","Ruby","Go"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
