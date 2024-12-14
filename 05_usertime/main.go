package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in GO")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// remember about the argument of Format - this is it
	fmt.Println(presentTime.Format("15:04:05 02-01-2006 Monday 3PM MST"))

	createdDate := time.Date(2024, time.January, 25, 15, 35, 0, 0, time.UTC)
	fmt.Println("createdDate:", createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday 15:04:05"))
}
