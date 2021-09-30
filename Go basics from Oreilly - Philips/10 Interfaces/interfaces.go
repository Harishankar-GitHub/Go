package main

import "fmt"

type BasePerson struct {
	First string
	Last  string
}

type Employee struct {
	BasePerson
	Salary int
	Boss   *Manager
}

type Manager struct {
	Employee
}

type Person interface {
	GetName() string
}

func SayHello(p Person) {
	fmt.Printf("Hello, %s\n", p.GetName())
}

func (e Employee) GetName() string {
	return e.First
}

func (m Manager) GetName() string {
	return m.First
}

func main() {
	m := &Manager{
		Employee{
			BasePerson: BasePerson{
				First: "Manager's First Name",
				Last:  "Manager's Last Name",
			},
			Salary: 40000,
			Boss:   nil,
		},
	}

	e := &Employee{
		BasePerson: BasePerson{
			First: "Employee's First Name",
			Last:  "Employee's Last Name",
		},
		Salary: 30000,
		Boss:   m,
	}

	fmt.Println(m.First)
	fmt.Println(e.First)

	SayHello(m)		// Calling SayHello() and that calls Manager's GetName() and Employee's GetName() methods respectively.
	SayHello(e)

	// ----------------------------------------------------------------------------------------------------------

	people := []Person{Person(m), Person(e)}
	// people := []Person{m, e}	// This also works!

	for _, person := range people {
		SayHello(person)
	}
}
