package main

import "fmt"

// Slices are variable length arrays for storing a single datatype.
// Maps are key value pairs.

func main() {
	s := []int{}
						// [] to indicate that this is a slice.
						// int{} to specify the type.
	fmt.Println(s)			// Prints []

	s1 := []int{1,2,3,4}		// Creating a slice with elements.
	fmt.Println("S1: ", s1)		// Prints [1 2 3 4]

	s1 = append(s1, 5)		// Appending 5 to the slice.
	fmt.Println("S1: ", s1)		// Prints [1 2 3 4 5]

	s2 := make([]int, 0)		// Creating a slice using make keyword. Length = 0, Datatype - int.
	fmt.Println("S2: ", s2)		// []

	s3 := make([]int, 10)		// Creating a slice using make keyword. Length = 3, Datatype - int.
	fmt.Println("S3: ", s3)		// Prints [0 0 0 0 0 0 0 0 0 0]
}