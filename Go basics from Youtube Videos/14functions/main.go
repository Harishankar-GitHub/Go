package main

import "fmt"

// main() function acts as an entrypoint for the program.
func main() {
	fmt.Println("Functions in Golang!")
	greeter()
	anotherGreeter()

	result := adder(3, 5)
	fmt.Println("Result from adder()", result)

	result = proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result from proAdder()", result)
}

func greeter() {
	fmt.Println("Hello, from Golang!")
}

func anotherGreeter() {
	fmt.Println("Hello again, from Golang!")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}
