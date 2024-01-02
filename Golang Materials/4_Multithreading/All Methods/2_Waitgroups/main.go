package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually this are pointers

func main() {
	websitelist := []string {
		"https://www.google.com",
		"https://www.meta.com",
	}

	var sum int
	sum = 0 

	for _, web := range websitelist{
		wg.Add(1)
		go getStatusCode(web, &sum)
	}

	wg.Wait()
	fmt.Println("Sum : ",sum)
}

func getStatusCode(endpoint string, sum *int)  {
	defer wg.Done()

	(*sum)++

	result, err:= http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Status Code : %d on web : %s\n",result.StatusCode, endpoint)
}