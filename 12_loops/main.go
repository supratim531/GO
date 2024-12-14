package main

import "fmt"

func main() {
	fmt.Println("This is about loops")

	// var days = [3]string{"en", "de"}
	var days []string
	fmt.Println("Size before:", len(days))

	days = append(days, "Sun", "Mon", "Tue", "Wed", "Thu")
	fmt.Println("Size after:", len(days))

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i, day := range days {
		fmt.Printf("%d: %s\n", i, day)
	}
}
