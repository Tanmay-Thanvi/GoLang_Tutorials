package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request in Golang")
	// http package is imp for that

	response,err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response : %v,\nType is %T\n",response,response)

	defer response.Body.Close() // caller's responsibility to close the connection 

	// Majority will be done by ioutil
	databytes,err := ioutil.ReadAll(response.Body)

	if err!=nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content : "+content)
}
