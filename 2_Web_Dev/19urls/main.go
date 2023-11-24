package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=yhsjk1212"

func main() {
	fmt.Println("Handling Urls in Golang")
	fmt.Println("Url : ", myurl)

	// parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println("Result : ", result.Scheme)
	fmt.Println("Result : ", result.Host)
	fmt.Println("Result : ", result.Path)
	fmt.Println("Result : ", result.RawQuery)
	fmt.Println("Result : ", result.Port())

	qparams := result.Query()
	fmt.Printf("Query Parameters : %v and its type : %T\n",qparams,qparams)

	fmt.Println(qparams["coursename"])

	for _,v := range(qparams){
		fmt.Println("Param is ",v)
	}

	// Creating url 

	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=tanmay",
	}

	anotherurl := partsofUrl.String()
	fmt.Println(anotherurl)
}
