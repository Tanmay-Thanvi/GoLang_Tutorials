package main

import (
	"errors"
	"fmt"
)

func main() {
	err := some_error_operation()

	// an only way to handle error in go (and log.fatal)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully performed operation")
	}
}

// Check try-catch block 

// Tip 1 : Errors.new()
func some_error_operation() error {
	return errors.New("Some error occured") // concrete error 
}

// Tip 2 : Use Error Interface and structs (better than just strings)
type My_Error_Struct struct {
	err_msg error // error string msg 
	err_code int
}

// Tip 3 : Return errors as early as possible (do not do it in nested ifs)

// Tip 4 : Errors.Is .As .Wrap (2 errors are under a same type)

// Tip 5 : Reuse common errors 
var global_lvl_err error = errors.New("Common Error") // Readability and Maintainability

// Tip6 : Utilizing Request scoped data 

// Log.Fatal is also another method for error handling 