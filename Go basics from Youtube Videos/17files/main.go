package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")

	content := "This needs to go in a file"
	// Content of the file.

	file, err := os.Create("./mygofile.txt")
	// Creating the file using the os package of Golang.
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	// Writing the content to the file using the io package of Golang.
	checkNilErr(err)

	fmt.Println("Length is", length)

	defer file.Close()
	// Closing the file. Used defer keyword here so that if any code is added
	// below this line, the file will be closed only at the end of the function.

	readFile("./mygofile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName) // returns []byte, error
	checkNilErr(err)

	fmt.Println("Content inside the file is", dataByte)
	fmt.Println("Content inside the file is", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
