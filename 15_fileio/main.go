package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling...")
	content := "This needs to go inside a file - LearnCodeOnline.in"

	file, err := os.Create("./lco-go-file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v bytes have written\n", length)
	}

	defer file.Close()
	readFile("./lco-go-file.txt")
}

func readFile(filePath string) {
	databyte, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The file contains:")
		fmt.Println(string(databyte))
	}
}
