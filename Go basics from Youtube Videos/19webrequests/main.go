package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://golang.org/"

func main() {
	fmt.Println("Handling Web Requests in Golang")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close()	// Caller's responsibility to close the connection.

	databytes, err := ioutil.ReadAll(response.Body)
	// Hover over ReadAll() to know more about it.
	checkNilErr(err)

	content := string(databytes)
	fmt.Println("Content", content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
