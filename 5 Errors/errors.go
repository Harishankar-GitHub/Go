package main

// In Go language, there's no Exception Handling.
// Instead, functions return an error variable that we need to check if there was an error.

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int

	for _, a := range os.Args[1:] {
		// os.Args[1:] - This will take from 1st parameter
		// os.Args[2:] - This will take from 2nd parameter
		i, err := strconv.Atoi(a)
				// Atoi is equivalent to ParseInt
				// Returns (int, error)

		if err != nil {
			panic(fmt.Sprintf("Invalid value: %v", err))	// panic will exit the program immediately.
		}
		sum += i
	}
	fmt.Printf("Sum: %v\n", sum)
}

// Running this program
// go run fileName.go param1 param2 param2 ...
// go run errors.go 1 2 3 4