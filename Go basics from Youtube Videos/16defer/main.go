package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	// Prints:
	// Hello
	// Two
	// One
	// World

	// The statement that has defer keyword is executed at the end of the function.
	// If we have multiple deferred statements, they are executed in the reverse (LIFO).

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
