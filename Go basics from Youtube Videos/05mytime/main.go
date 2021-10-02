package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time package in Golang")

	presentTime := time.Now()
	fmt.Println("Time:", presentTime)

	fmt.Println("Formatted time:", presentTime.Format("01-02-2006"))
	// 01-02-2006 must be used to specify the format.
	fmt.Println("Formatted time:", presentTime.Format("01-02-2006 Monday"))
	// Monday must be used if we need to print Day.
	fmt.Println("Formatted time:", presentTime.Format("01-02-2006 15:04:05 Monday"))
	// 15:04:05 must be used if we need to print hh:mm:ss.

	createdDate := time.Date(2020, time.July, 10, 17, 30, 05, 0, time.UTC)
	fmt.Println("Created Date:", createdDate)
	fmt.Println("Formatted Created Date", createdDate.Format("01-02-2006 Monday"))

	fmt.Println("Time - Nanosecond:", time.Now().Nanosecond())	// Prints in Nanosecond
	fmt.Println("Time - UnixNano:", time.Now().UnixNano())		// Prints more accurately and has more digits.
}
