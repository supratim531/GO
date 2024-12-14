package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is hilarious")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Enter your age: ")

	if age, err := strconv.ParseUint(strings.TrimSpace(input), 0, 8); age >= 40 {
		if err != nil {
			fmt.Printf("Jaahhh.!! Eeee ki holo... %v \n", err)
		} else {
			fmt.Println("Dhere Buro")
		}
	} else {
		if err != nil {
			fmt.Printf("Jaahhh.!! Eeee ki holo... %v \n", err)
		} else {
			fmt.Println("Hustle")
		}
	}
}
