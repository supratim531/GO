package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/supratim531/gomongoapi/router"
)

func main() {
	fmt.Println("This is utilizing the MongoDB")

	router := router.Router()

	fmt.Println("Listening at port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
