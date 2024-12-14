package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://sayan-textutils.netlify.app"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is %T\n", resp)
	fmt.Printf("response body type is %T\n", resp.Body)
	fmt.Printf("response header type is %T\n", resp.Header)

	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
