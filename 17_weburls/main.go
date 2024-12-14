package main

import (
	"fmt"
	"net/url"
)

const hecticUrl = "https://lco.dev:3000/course/learn?course_name=reactjs&payment_status=pending"
const hecticUrl2 = "https://lco.dev:3000/course/learn?course_name=reactjs,java&payment_status=pending"
const hecticUrl3 = "https://lco.dev:3000/course/learn?course_name=reactjs&course_name=java&payment_status=pending"

func main() {
	fmt.Println("Now playing with URLs...")

	result, err := url.Parse(hecticUrl3)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T, %v\n", result.Scheme, result.Scheme)
	fmt.Printf("%T, %v\n", result.Host, result.Host)
	fmt.Printf("%T, %v\n", result.Path, result.Path)
	fmt.Printf("%T, %v\n", result.RawQuery, result.RawQuery)
	fmt.Printf("%T, %v\n", result.Port(), result.Port())

	var queryParams = result.Query()
	fmt.Printf("%T, %v\n", queryParams, queryParams)

	for _, value := range queryParams {
		fmt.Println("The value is", value)
	}

	customUrlParts := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn/course",
		RawQuery: "user=hitesh",
	}

	customUrl := customUrlParts.String()
	fmt.Println("This is your custom URL:", customUrl)
}
