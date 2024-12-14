package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello")

	var fruits = []string{"Apple", "Orange"}
	fmt.Println(fruits)
	fmt.Printf("The type of fruits is %T\n", fruits)

	fruits = append(fruits, "Mango", "Banana")
	fmt.Println(fruits)
	fmt.Printf("The type of fruits is %T\n", fruits)
	fmt.Println("The type of fruits is", fruits[:2])

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 935
	highScores[2] = 136
	highScores[3] = 837

	fmt.Printf("%T\n", highScores)
	fmt.Println(highScores, len(highScores))

	highScores = append(highScores, 555, 666, 321)
	fmt.Printf("%T\n", highScores)
	fmt.Println(highScores, len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores, len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	var stores []string
	fmt.Println(stores, len(stores))
	stores = append(stores, "ABC Store", "DEF Store")
	fmt.Println(stores, len(stores))

	var removeIndex int = len(highScores) - 1
	highScores = append(highScores[:removeIndex], highScores[removeIndex+1:]...)
	fmt.Println("New highScores:", highScores)
}
