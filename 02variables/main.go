package main

import "fmt"

const LoginToken string = "a;lsfhi" // public variable
const loginToken string = "asfdasd" // private variable

func main() {
	fmt.Println("Variables")

	var username string = "Abhishek Jaglan"

	fmt.Println(username)

	fmt.Printf("Variable is of type: %T\n", username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type: %T\n", isTrue)

	var smallVal uint16 = 2546
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFLoat float32 = 2546.7865875785
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T\n", smallFLoat)

	var bigFloat float64 = 2546.7865875785
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T\n", bigFloat)

	// default values and some aliases
	var someVar int
	fmt.Println(someVar) // value is 0, no garbage value
	fmt.Printf("Variable is of type: %T\n", someVar)

	var someStr string
	fmt.Println(someStr) // blank string
	fmt.Printf("Variable is of type: %T\n", someStr)

	var someFloat float64
	fmt.Println(someFloat) //  value is 0, no garbage value
	fmt.Printf("Variable is of type: %T\n", someFloat)

	// implicit type of declaring var

	var website = "abhishekjaglan.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	// no var usage
	// walrous operator allowed only inside methods
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	// public variable access
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)

	fmt.Println(loginToken)
}
