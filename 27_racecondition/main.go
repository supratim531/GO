package main

import (
	"fmt"
	"sync"
)

func main() {
	// RaceCondition()
	ResolvedRaceCondition()
}

func RaceCondition() {
	var scores = []int{0}
	var wg = &sync.WaitGroup{}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		scores = append(scores, 1)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		scores = append(scores, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		scores = append(scores, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(scores)
}

func ResolvedRaceCondition() {
	var scores = []int{0}
	var wg = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("One R")
		mutex.Lock()
		scores = append(scores, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Two R")
		mutex.Lock()
		scores = append(scores, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Three R")
		mutex.Lock()
		scores = append(scores, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(scores)
}
