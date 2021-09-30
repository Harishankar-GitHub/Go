package main

// Types such as structures can have method sets.
// This means that the type has functions that can be called on an instance of the type.
// This is similar to how classes work in other languages.

// Lower case method names are private to the package.
// Upper case method names are public and can be accessed outside the package.

import (
	"fmt"
	"time"
)

type Person struct {
	First string
	Last string
	Dob time.Time
}

func NewPerson(first, last string, year, month, day int) *Person {
	// Return value is a Pointer to a Person.
	// This method returns a new Person along with the DOB.

	return &Person{
		First: first,
		Last: last,
		Dob: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

// GetAge() is a method that belongs to the Person type.
func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365)
}

// SayHello() is a method that belongs to the Person type.
func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.First)
}

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
	fmt.Println("Hello,", name)
}

func getHello(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}

func main() {
	sayHello("Harish")

	s := getHello("Harish")
	fmt.Println(s)

	p := &Person{
		First: "Harishankar",
		Last: "Bhat R",
	}
	p.SayHello()

	q := NewPerson("Harishankar", "Bhat R", 1997, 07, 06)
	q.SayHello()
	fmt.Println("Age: ", q.GetAge())	
}