package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	harish := User{"Harish", "harish@mail", true, 24}
	fmt.Println(harish)

	fmt.Printf("Harish details are: %+v\n", harish)
	// %+v gives more details of the struct.

	fmt.Printf("Name is %v and Email is %v.\n", harish.Name, harish.Email)

	harish.GetStatus()
	harish.NewMail()

	fmt.Printf("Name is %v and Email is %v.\n", harish.Name, harish.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Normal functions are called as functions.
// A function that belongs to any Struct is called as a Method.
// The below methods belongs to User Struct. So it is called as a Method instead of a Function.
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	// This doesn't change the actual User Object.
	// A copy of User Object is sent here.
	u.Email = "test@email.com"
	fmt.Println("Email of this user is:", u.Email)
}
