package main

import (
	"fmt"
)

type User struct {
	Name   string
	Status bool
	Age    uint8
}

func main() {
	fmt.Println("Welcome to struct")

	var sayan User = User{
		Name:   "John",
		Status: true,
		Age:    30,
	}
	supratim := User{"Brock", true, 23}

	fmt.Println(sayan)
	fmt.Println(supratim)
	fmt.Printf("details of sayan: %+v\n", sayan)
	fmt.Printf("details of supratim: %+v\n", supratim)
	fmt.Printf("Name: %s\nStatus: %t\nAge: %d\n", supratim.Name, supratim.Status, supratim.Age)
}
