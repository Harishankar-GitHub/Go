package main

import "fmt"

// Structures provide a way to combine data with different types into a single data structure.

type Person struct {
	First string
	Last string
	Age int
	Phone Phone		// Both `Phone Phone` and `Phone` works.
}

type Phone struct {
	AreaCode string
	Prefix string
	Suffix string
}

func main() {
	p := Person {
		First: "Harishankar",
		Last: "Bhat R",
		Age: 23,
		Phone: Phone{
			AreaCode: "123",
			Prefix: "Prefix",
			Suffix: "Suffix",
		},				// Last field must have a coma. If not, Go language won't compile the code.
	}

	q := &Person{"Harish", "Jack", 24, Phone{AreaCode: "123", Prefix: "456", Suffix: "789",}}
						// Assigning the address of Person to q with different values.

	r := Person{"Harish", "Jack", 24, Phone{AreaCode: "456", Prefix: "789", Suffix: "123",}}
						// Creating a new Structure and assigning it to r.

	fmt.Println("P: ", p);
	fmt.Println("Q: ", q);
	fmt.Println("R: ", r);
}