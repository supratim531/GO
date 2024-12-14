package main

import "fmt"

func main() {
	fmt.Println("This is about pointers")

	a := 25
	var ptr *int = &a
	*ptr = *ptr + 12
	fmt.Println("The value of a is", a)

	var fruitList [2]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}
