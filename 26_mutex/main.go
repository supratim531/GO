package main

import (
	"fmt"
	"net/http"
	"sync"
)

// usually we keep these as pointers
var mutex sync.Mutex
var wg sync.WaitGroup

// dummy signals to test lock-unlock
var signals = []string{"test"}

func main() {
	fmt.Println("Wait Groups!")

	var websiteList = []string{
		"https://lco.dev",
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://fb.com",
		"https://go.dev",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

	// mimicMain()
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Ooops in %v\n", endpoint)
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()

		fmt.Printf("%d OK for %v\n", res.StatusCode, endpoint)
	}
}

// func mimicMain() {
// 	fmt.Println("Wait Groups!")

// 	var websiteList = []string{
// 		"https://lco.dev",
// 		"https://google.com",
// 		"https://facebook.com",
// 		"https://amazon.com",
// 		"https://fb.com",
// 		"https://go.dev",
// 	}

// 	for _, website := range websiteList {
// 		mimicGetStatusCode(website)
// 	}
// }

// func mimicGetStatusCode(endpoint string) {
// 	res, err := http.Get(endpoint)

// 	if err != nil {
// 		fmt.Printf("Ooops in %v\n", endpoint)
// 	} else {
// 		fmt.Printf("%d OK for %v\n", res.StatusCode, endpoint)
// 	}
// }
