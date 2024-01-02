package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Sender goroutine
	go send(&ch)

	// Receiver goroutine
	receive(&ch)
}

func send(ch *chan int) {
	var sample int = 50
	*ch <- 5 // Sending value into the channel
	*ch <- 10
	*ch <- sample
	close(*ch)
}

func receive(ch *chan int) {
	// time.Sleep(5 * time.Second) // Point of discussion 
	for value := range *ch {
		fmt.Println(value) // Receiving values from the channel
	}
}
