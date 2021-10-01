package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv.ParseFloat() accepts a string and bitSize.
	// For instance, if the input is 5, it is not only 5.
	// After entering 5, we press enter to pass the input to the program.
	// So along with 5, we also get a \n.
	// Now the input looks like 5\n.
	// That is the reason we make use of strings package to trim the space.
	// In this case, the bitSize for the float is 64.

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
