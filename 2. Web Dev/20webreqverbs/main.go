package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get & Post Request in Golang")

	// PerformGetRequest() 
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest()  {
	const myurl = "http://localhost:8000/get"

	response,err := http.Get(myurl)

	CheckError(err)

	defer response.Body.Close()

	fmt.Println("Response : ",response)
	fmt.Println("Status Code : ",response.StatusCode)
	fmt.Println("Content Length : ",response.ContentLength)

	var responseString strings.Builder
	databytes,err := ioutil.ReadAll(response.Body) // This is an inevitable line
	byteCount, _ := responseString.Write(databytes)

	fmt.Println("Byte Cnt is ",byteCount)
	fmt.Println(responseString.String())
	// fmt.Println("Content : ",string(databytes))
}

func PerformPostJsonRequest()  {
	const myurl = "http://localhost:8000/post"

	// fake json payload 

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with Golang",
			"price" : 0,
			"platform" : "Youtube"
		}
	`)

	response , err := http.Post(myurl,"application/json",requestBody)
	CheckError(err)

	defer response.Body.Close();

	databytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content : ",string(databytes))
}

func PerformPostFormRequest()  {
	const myurl = "http://localhost:8000/postform"

	// form data 

	data := url.Values{}
	data.Add("firstname","Tanmay")
	data.Add("lastname","Thanvi")
	data.Add("email","tnt@gmail.com")

	response ,err := http.PostForm(myurl,data)
	CheckError(err)
	defer response.Body.Close()
	databytes , _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content : ",string(databytes))
}

func CheckError(err error)  {
	if err != nil {
		panic(err)
	}
}