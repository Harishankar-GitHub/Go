package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://test.dev:3000/learn?coursename=reactjs&paymentid=ghuy456"

func main() {
	fmt.Println("Welcome to URLs in Golang")
	fmt.Println(myUrl)

	// Parsing the URL
	result, _ := url.Parse(myUrl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of Query Params is %T\n", queryParams)

	fmt.Println(queryParams["coursename"])

	for _, val := range queryParams {
		fmt.Println("Param is", val)
	}

	// Creating the URL
	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "Harishankar-GitHub",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
	// Ctrl + Click the URL in the console to open in browser.
}
