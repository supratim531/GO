package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is all about channels in goroutines")

	myCh := make(chan int, 3)
	var wg = &sync.WaitGroup{}

	wg.Add(2)
	// myCh <- 5
	// fmt.Println(<-myCh)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(ch) // cannot close in recieve-only channel

		fmt.Println("First 1")
		fmt.Println("From First 1: Hey recieved 1st:", <-ch)
		fmt.Println("From First 1: Hey recieved 2nd:", <-ch)

		defer wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Second 1")
		ch <- 108
		ch <- 149
		ch <- 111
		ch <- 115
		close(ch)

		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}
