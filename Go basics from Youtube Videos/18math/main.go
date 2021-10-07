package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Math and Random Number in Golang")

	var numberOne int = 2
	var numberTwo float64 = 4.5

	fmt.Println("The sum is", numberOne+int(numberTwo))

	// Random Number
	rand.Seed(time.Now().UnixNano())
	// If the above line is not present, we get same number generated everytime.
	fmt.Println(rand.Intn(5) + 1)
	// This generates random number between 1 to 6.
}
