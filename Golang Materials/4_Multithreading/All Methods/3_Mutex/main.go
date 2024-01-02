package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Problem on waitgroups occurs when 3 or 4 go routines are fired up simaltaneously that share same memory 
var signals = []string{"test"}

var wg sync.WaitGroup // usually this are pointers
var m sync.Mutex

func main() {
	websitelist := []string {
		"https://www.google.com",
		"https://www.meta.com",
	}

	for _, web := range websitelist{
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()

	fmt.Println("Signals : ",signals)
}

func getStatusCode(endpoint string)  {
	defer wg.Done()

	result, err:= http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	m.Lock() // requirement of mutex
	signals = append(signals, endpoint)
	m.Unlock()

	fmt.Printf("Status Code : %d on web : %s\n",result.StatusCode, endpoint)
}