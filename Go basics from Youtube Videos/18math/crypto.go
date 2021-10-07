package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Crypto in Golang")

	// Random Number using Crypto.
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("My random number", myRandomNumber)
}
