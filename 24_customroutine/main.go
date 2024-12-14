package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter2("Hello")
	greeter("World")
}

func greeter2(s string) {
	for i := 0; i < 6; i++ {
		// time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func greeter(s string) {
	time.Sleep(1 * time.Millisecond)

	for i := 0; i < 6; i++ {
		// time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
