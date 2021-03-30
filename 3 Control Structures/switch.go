package main

import "fmt"

func main() {
	a := 20

	switch a == 10 {
	case true:
		fmt.Println("Yes!")
	case false:
		fmt.Println("No!")
	default:
		fmt.Println("None of these!")
	}

	switch a {
	case 10:
		fmt.Println("It is 10!")
	case 5:
		fmt.Println("It is 5!")
	default:
		fmt.Println("None of these!")
	}
}