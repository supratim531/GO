package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is like match case")

	rand.Seed(int64(time.Now().Nanosecond()))
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("This is 1! You got 2!")
		fallthrough
	case 2:
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	case 4:
		fmt.Println("This is 4")
	case 5:
		fmt.Println("This is 5")
	case 6:
		fmt.Println("This is 6! Play again")
		fallthrough
	default:
		fmt.Println("What was that!")
	}
}
