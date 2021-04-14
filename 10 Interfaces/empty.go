package main

import "fmt"

// In Go, Interfaces provide an abstraction so that different types that implement the same methods
// can be treated the same way.

// Every type in Go, implements empty interface.

func main() {
	var x interface{}		// We can assign anything to this interface.
	x = "Hello world!"
	fmt.Println(x)		// Printing the value.

	var y interface{}
	y = 100
	fmt.Println(y)		// Printing the value.


	var z interface{}
	z = "Hello world!"
	// s := z.(string)	// We are specifying the type of the data z.(string) that is stored in the Interface. s := z also works.
					// What if we don't know the type ?
	s, ok := z.(string)	// A 2nd return argument can be specified.
	if !ok {
		panic("NO!")
		// To print NO, make z = 10.
		// Or, make s, ok := z.(int)
	}
	fmt.Println(s)


	var a interface{}
	a = 100

	switch a.(type) {			// Checking the type using Switch.
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String!")
	}
}