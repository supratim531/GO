package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Requesting my AWS...")
	PerformGetRequest("http://35.200.196.220/api/v1/doctor")
}

func PerformGetRequest(customUrl string) {
	resp, err := http.Get(customUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Content length:", resp.ContentLength)

	databyte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v -> %T\n", string(databyte), databyte)

	var responseString strings.Builder

	bytesCount, err := responseString.Write(databyte)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseString)
	fmt.Println("Also get the Content length in this way:", bytesCount)
	fmt.Println("Also get the JSON in this way:", responseString.String())
}
