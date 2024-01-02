package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Math, Crypto, Random Number")

	var my_num1 int = 2
	var my_num2 float64 = 4.5

	fmt.Println("Sum : ", my_num1+int(my_num2))

	// random number - 2 ways (math & crypto)

	// math 
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random no. : ", rand.Intn(5)+1)

	// crypto 
	random_no, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(random_no)
}
