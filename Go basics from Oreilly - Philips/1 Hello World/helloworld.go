package main
// "package main" is used to write an executable program.
// If this file is going to be used anywhere else, we can use "package anyOtherName".

import "fmt"
// This is a standard library (Format).
// Format is a package that includes a function for writing the characters to the screen.

// Declaring a function:
// 	Syntax - func functionName() {}
// 	In this case, this is an executable program. So, we use "func main() {}".
//	"func main() {}" is the entry point for the Go program and it doesn't take any arguments.
func main() {
	fmt.Println("Hello world!")
}

// Running the program
// 2 options:
//	1. go run fileName.go 		// Creates a temporary binary that will be executed.
//	2. go build fileName.go		// Creates a separate executable file with the same name.
//	   ./fileName			// To run the separate executable file.