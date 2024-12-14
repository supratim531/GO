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
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learn how to create a JSON from a struct/interface")

	course1 := course{
		Name:     "ReactJS",
		Price:    1500,
		Platform: "Udemy",
		Password: "password123",
		Tags:     nil,
	}

	// course1.EncodeCourseToJSON()

	var courses []course = []course{
		{Name: "Go", Price: 1900, Platform: "Udemy", Password: "go123", Tags: []string{"Go", "Backend"}},
		{Name: "Python", Price: 1800, Platform: "Udemy", Password: "python123", Tags: []string{"Python", "Backend"}},
		{Name: "Node.js", Price: 1600, Platform: "Udemy", Password: "node123", Tags: []string{"JavaScript", "Node.js", "Backend"}},
	}

	multipleCoursesData := map[string]interface{}{
		"message": "Welcome Sir",
		"data":    courses,
	}

	oneCourse, _ := EncodeToJSON(course1)
	multipleCourses, _ := EncodeToJSON(multipleCoursesData)

	fmt.Printf("%s -> %T\n", oneCourse, oneCourse)
	fmt.Printf("%s -> %T\n", multipleCourses, multipleCourses)
}

func (c course) EncodeCourseToJSON() int {
	fmt.Printf("%+v -> %T\n", c, c)

	// jsonData, err := json.Marshal(c)
	jsonData, err := json.MarshalIndent(c, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	} else {
		fmt.Println("JSON data:", jsonData)
		fmt.Printf("%s -> %T\n", jsonData, jsonData)
	}

	return 108
}

func EncodeToJSON(data interface{}) ([]byte, error) {
	// jsonData, err := json.Marshal(data)
	jsonData, err := json.MarshalIndent(data, "->", "  ")
	return jsonData, err
}
