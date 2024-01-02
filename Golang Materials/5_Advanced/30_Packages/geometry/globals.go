package geometry

import (
	"fmt"
	"math"
)

var Pi float64 = math.Pi
var e float64 = math.E

func PrintPi() { // Capital is must
	fmt.Println("Value of E (private member) : ",e)
	fmt.Println("Value of Pi (public member) : ", Pi)
}
