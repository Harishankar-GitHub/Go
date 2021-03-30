package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Tick")
		time.Sleep(1 * time.Second)
		break	// If break is not given, Tick will be printed every second until we interrupt (Ctrl+C).
	}
}