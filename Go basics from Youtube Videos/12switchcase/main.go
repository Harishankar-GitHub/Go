package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of Dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open!")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
		// If we hit case 3, the flow will go to case 4 as well.
		// And from case 4 it goes to case 5 because case 4 also has fallthrough.
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
		// If we hit case 4, the flow will go to case 5 as well.
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again!")
	default:
		fmt.Println("What was that!")
	}
}
