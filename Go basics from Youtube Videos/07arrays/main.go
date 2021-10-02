package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	var fruitList [4]string
//	var arrayName [length]type
//	If we use array, length of the array must be specified.
	
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is:", fruitList)
	// Prints [Apple Tomato  Orange]
	// Index 2 has a space.
	// It is an indication that there's some empty space in the Array.

	fmt.Println("Length of Fruit List is:", len(fruitList))
	// Prints 4.
	// No matter how many values are in the Array, it always prints the length that
	// was specified at the beginnning.

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veggie List is:", vegList)
	fmt.Println("Length of Veggie List is:", len(vegList))
}
