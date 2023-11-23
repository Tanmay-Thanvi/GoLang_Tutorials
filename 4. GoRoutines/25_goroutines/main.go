package main

import (
	"fmt"
	"net/http"
	"sync"
)

// "fmt"
// "time"

var wg sync.WaitGroup // this are pointers 
var mysignals = []string{"test"}

func main() {
	// go greeter("Hello") // go routine created
	// greeter("World")

	// if we dont use sleep then just world greeter gets printed and main method gets completed
	// when thread of hello comes till that main function is over
	// eg :- u involved in tv so much that u forget to start ac and eat at the same time (parllelism example)

	// sync package is prefered for this purpose than sleep

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	usegokeyword := 1

	for _, web := range websitelist {
		if usegokeyword == 1 {
			// this will not give any output bcoz it is not waiting for response 
			go getStatusCode(web)
			wg.Add(1) // go for firing one routine only 
		} else {
			getStatusCode(web)
		}
	}

	wg.Wait() // responsible for not execution of main method 
	fmt.Println(mysignals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) // 1st hack
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops in endpoint : ",endpoint)
	} else {
		mysignals = append(mysignals, endpoint)
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
