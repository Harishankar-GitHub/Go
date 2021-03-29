package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I'm a String"
	fmt.Printf("Ends with String ? %v\n", strings.HasSuffix(s, "String"))
}