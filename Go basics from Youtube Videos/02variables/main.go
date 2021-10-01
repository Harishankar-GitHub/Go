package main

import "fmt"

// Outside a method or global stuff, this type of declaration is:

// Allowed:
// var token_1 = 10
// var token_2 int = 20

// Not allowed:
// token_3 := 30

// Declaring Constants
const LoginToken string = "fdsakoiqwzclmnbqw"
// The 1st letter of the variable LoginToken is capitalized.
// In Golang, to make public, 1st letter must be capitalized.

func main() {

	// %T - a Go-syntax representation of the type of the value

	var username string = "Harish"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	// uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455454454554455454455
	// 5 values after the decimal are considered. Rest of them are ignored.
	var bigFloat float64 = 255.455454454554455454455
	// More precise value compared to float32 is printed.
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)


	// Default values and some Aliases

	var anInt int			// Default value of int is 0
	fmt.Println(anInt)
	fmt.Printf("Variable is of type: %T \n", anInt)

	var aString string		// Default value of string is empty
	fmt.Println(aString)
	fmt.Printf("Variable is of type: %T \n", aString)


	// Implicit Type

	var website = "www.google.com"
	// We haven't specified the type here.
	// Here comes the role of Lexer.
	// It will automatically interpret the type of the variable.
	fmt.Println(website)


	// No var style

	numberOfUsers := 300000
	// var keyword and type specification is not required.
	// If we change the value to 300000.4554, it accepts.
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
