package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// We have a new reader.
	// From where do we read ?
	// For that we have a libary called "os".
	// From os, we want to read from Standard input and output.

	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || comma err ok syntax

	// In Golang, there's no try-catch concept.
	// The language design of Go expects to treat problems or errors like variables.

	input, _ := reader.ReadString('\n')
	// We are gonna keep on reading till we get a \n.
	// The value that we read will be treated as string.
	// Here the _ represents the error.
	// If we want to store the error we can use input, err
	// If we don't care about the error, we can use input, _

	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}