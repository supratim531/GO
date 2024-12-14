package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")
	defer deferFunc()
}

func deferFunc() {
	defer fmt.Println()

	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
