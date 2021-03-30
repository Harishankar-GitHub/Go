package main

import "fmt"

func main() {
	a := 5

	if a < 0 {
		fmt.Printf("Your value is negative!")
	} else if a < 10 {
		fmt.Printf("Your value is single digit!")
	} else {
		fmt.Printf("Your value has multiple digits")
	}
}