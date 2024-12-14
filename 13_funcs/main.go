package main

import (
	"fmt"
	"sort"
)

func main() {
	sum, def := proAdder(542, 65, 2, 867, 2424)
	fmt.Println("Lets talk about functions")
	fmt.Println("Here is the sum", sum, def)

	var supratim User = User{"Sayan", "Das", 56, false}
	fmt.Println("Status:", supratim.status)
	supratim.GetStatus()

	supratim.SetStatus(true)
	supratim.GetStatus()
}

func proAdder(values ...int) (int, string) {
	sum := 0
	fmt.Printf("Type of values is %T \n", values)

	fmt.Println("before sort", values)
	sort.Ints(values)
	fmt.Println("after sort", values)

	for _, val := range values {
		sum += val
	}

	return sum, "Default Message"
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	status    bool
}

func (u User) GetStatus() {
	fmt.Printf("The status of %v is %v\n", u.FirstName, u.status)
}

func (u User) SetStatus(status bool) {
	u.status = status
	fmt.Println(u.status)
}
