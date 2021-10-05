package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ { // There is no ++d syntax in Golang!
		fmt.Println(days[d])
	}

	for i := range days {
		// In Golang, this i returns the index, not the actual value.
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}

	for _, value := range days {
		fmt.Printf("Value is %v\n", value)
	}

	someValue := 1

	for someValue < 10 {
		if someValue == 5 {
			break
		}
		fmt.Println("Value is", someValue)
		someValue++
	}

	someValue = 1
	for someValue < 10 {
		if someValue == 5 {
			someValue++
			continue
		}
		fmt.Println("Value is", someValue)
		someValue++
	}

	someValue = 1
	for someValue < 10 {
		if someValue == 2 {
			goto someLabel
		}
		fmt.Println("Value is", someValue)
		someValue++
	}

	someLabel:
		fmt.Println("Jumping at someLabel!")
}
