package main

import "fmt"

func main() {
	// Creating a map: Method 1
	map1 := map[string]string{}
	// Syntax = map1 := map[keyDatatype]valueDatatype{}

	map1["firstName"] = "Harishankar"
	map1["lastName"] = "Bhat R"

	fmt.Println("Map 1 First Name: ", map1["firstName"])
	fmt.Println("Map 1 Last Name: ", map1["lastName"])

	// Creating a map: Method 2
	map2 := map[string]string{
		"firstName": "Harishankar",
		"lastName": "Bhat R",
	}

	fmt.Println("Map 2 First Name: ", map2["firstName"])
	fmt.Println("Map 2 Last Name: ", map2["lastName"])

	// Creating a map: Method 3
	map3 := make(map[string]string)

	map3["firstName"] = "Harishankar"
	map3["lastName"] = "Bhat R"

	fmt.Println("Map 3 First Name: ", map3["firstName"])
	fmt.Println("Map 3 Last Name: ", map3["lastName"])
}