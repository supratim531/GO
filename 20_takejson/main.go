package main

import (
	"encoding/json"
	"fmt"
)

// It is started with small cap so this is not publicly available to other files
type course struct {
	Name     string   `json:"courseName"`
	Price    uint16   `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"password"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	acceptedJSONData := []byte(`{
		"courseName": "ReactJS",
		"price": 1500,
		"website": "Udemy",
		"password": "password@123",
		"tags": null
	}`)
	fmt.Printf("%v -> %T\n", acceptedJSONData, acceptedJSONData)

	acceptedJSONDataList := []byte(`{"data": [{
		"courseName": "ReactJS",
		"price": 1500,
		"website": "Udemy",
		"password": "password@123",
		"tags": null
	},
	{
		"courseName": "NodeJS",
		"price": 700,
		"website": "Udemy",
		"password": "password@123",
		"tags": [
		  "JavaScript",
      "NodeJS"
		]
	}]}`)
	fmt.Printf("%v -> %T\n", acceptedJSONDataList, acceptedJSONDataList)

	var course1 course
	var multipleCourses map[string][]course

	isValid := json.Valid(acceptedJSONData)

	if isValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(acceptedJSONData, &course1)
		fmt.Printf("%#v -> %T\n", course1, course1)
	} else {
		fmt.Println("JSON was invalid")
	}

	isValid = json.Valid(acceptedJSONDataList)

	if isValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(acceptedJSONDataList, &multipleCourses)
		fmt.Printf("%+v -> %T\n", multipleCourses, multipleCourses)
	} else {
		fmt.Println("JSON was invalid")
	}
}
