package main

import "fmt"

// First letter in capital shows that it is a public variable
const LoginToken string = "2eff8670nsfr2r"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("username is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("smallValue is of type: %T \n", smallValue)

	var smallFloatValue float64 = 255.4554628747546
	fmt.Println(smallFloatValue)
	fmt.Printf("smallFloatValue is of type: %T \n", smallFloatValue)

	// default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("website is of type: %T \n", website)

	// no var style
	numberOfUser := 100.12
	fmt.Println(numberOfUser)
	fmt.Printf("numberOfUser is of type: %T \n", numberOfUser)

	// public variable
	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type: %T \n", LoginToken)
}
