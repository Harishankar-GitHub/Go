package main

// Command Line Flags

// Go has a standard library for parsing flags on the command line.
// It is useful for small projects that don't need to conform to Unix standards.

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Flag package has lots of functions that help you to create command line arguments
// to the program.

var name string
var wait time.Duration
var debug bool

func main() {
	if name == "" {
		fmt.Println("Must add name to use this tool!")
		flag.Usage()	// Prints the usage of the Command Line Flag.
		os.Exit(1)
	}

	if debug {
		fmt.Printf("Going to wait for %v\n", wait)
	}
	time.Sleep(wait)
	fmt.Printf("Hello, %s\n", name)
}

// init() function can be placed in any of the Go files.
// The content of the init() function is executed prior to the program execution.
func init() {
	flag.BoolVar(&debug, "debug", false, "Turn on debugging output")
	flag.StringVar(&name, "name", "", "The name to say Hello to")
	
	defaultWait, err := time.ParseDuration("5s")
	if err != nil {
		panic("Could not parse default wait time")
	}
	flag.DurationVar(&wait, "wait-time", defaultWait, "Time to wait before print")
	flag.Parse()
}

// Running this file:
// 5 ways

// 1. go run flags.go	// In this method, the file name won't be printed properly.

// 2. go build flags.go	// The file name is printed properly.
//	ls
// 	./flags

// 3. ./flags --help OR ./flags -help

// 4. go run flags.go -name Harish
//	OR ./flags -name Harish
// 	OR ./flags -name Harish -wait-time 1s

// 5. ./flags -debug -name Harish
//	OR ./flags -debug -name Harish -wait-time 1s
