package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	err := some_operation(5)
	if err != nil {
		fmt.Println("Error ! :- ",err)
	}
}

// Tip 1 : Use errors.New()
func some_operation(a int) error {
	if a > 1 {
		return errors.New("Error Occured!")
	}
	return nil
}

// Tip 2 : Use error interface 
type error_string interface {

}

// Tip 3 : Return errors as early as possible 
