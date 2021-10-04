package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// No Inheritance in Golang.
	// No parent or super concepts in Golang.

	harish := User{"Harish", "harish@mail", true, 24}
	fmt.Println(harish)

	fmt.Printf("Harish details are: %+v\n", harish)
	// %+v gives more details of the struct.

	fmt.Printf("Name is %v and Email is %v.", harish.Name, harish.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// User, Name, Email, Status, Age - All these words have 1st letter capitalized
// because they are public.
