package main

import (
	"fmt"
	"sync"
)

func main() {
	// channels are the way using which multiple go routines talk to each other

	fmt.Println("Channels in Golang")

	myCh := make(chan int)
	wg := & sync.WaitGroup{}

	// fmt.Println(<-myCh) // This is deadlock
	// myCh <- 5 // this arrow is always like this

	wg.Add(2)

	// Recieve Only 
	go func (ch <-chan int, wg *sync.WaitGroup)  {
		// close(myCh) // will give error as channel is recieve only 

		// To check whether we are reading from close channel or channel actuaaly contains zero 
		val, isChannelOpen := <-myCh

		fmt.Println("Is Channel Open : ",isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	// Send Only 
	go func (ch chan<- int, wg *sync.WaitGroup)  {
		myCh <- 0
		close(myCh) // listening close channel - no error
		// myCh <- 5
		// myCh <- 6 // gives error there should be 2 listener 
		// close(myCh) // no issues 
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
