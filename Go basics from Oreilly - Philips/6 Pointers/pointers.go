package main

import "fmt"

func main() {
	s := "My String"
	sptr := &s

	fmt.Println("Original String: ", s)				// Original String.
	fmt.Println("Memory address: ", sptr)			// Prints the memory address.
	fmt.Println("String from memory address: ", *sptr)	// Prints the string from memory address.

	s2 := new(string)							// Creating a new String. Memory address will be assigned to s2.
	fmt.Println("s2 memory address: ", s2)			// Prints the memory address.
	fmt.Println("Value of s2 from memory address: ", *s2)	// Prints nothing because there's nothing in the memory.

	var s3 string
	var s4 *string							// s4 is a pointer to a String.

	fmt.Println("S3: ", s3)						// Prints nothing because there's no value.
	fmt.Println("S4: ", s4)						// Prints <nil> because there's no memory address assigned.

	s5 := "A String"
	var sptr1 *string
	fmt.Println("sptr1 value : ", sptr1)
	sptr1 = &s5								// := is required only at the time of assigning a value. While reassigning = is enough.

	fmt.Println("S5 value: ", s5)
	fmt.Println("S5 memory address: ", sptr1)

	i := new(int)
	fmt.Println("i memory address: ", i)
	fmt.Println("i value: ", *i)
}