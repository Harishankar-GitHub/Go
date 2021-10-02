package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int
	// * is the symbol that this datatype is of Pointer.
	// A Pointer which is going to be responsible for holding integers into that.
	// Here Pointer is just created.

	fmt.Println("Value of ptr is:", ptr)
	// When a pointer is declared and no value is assigned to it,
	// the default value of that is <nil>.

	myNumber := 25
	var myNumberPtr = &myNumber
	// Here Pointer is created which is also referencing memory.

	fmt.Println("Value of myNumberPtr is", myNumberPtr)
	// Prints the memory address.
	fmt.Println("Value of myNumberPtr is", *myNumberPtr)
	// Prints the actual value from the memory address.

	*myNumberPtr = *myNumberPtr * 2
	fmt.Println("Value after multiplication", myNumber)
	// Prints 50 because the operation is performed on the actual value.
}
