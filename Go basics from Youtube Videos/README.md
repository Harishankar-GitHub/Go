## Basics of Golang

> **IDE - Visual Studio Code**

### 01 Hello

- Created main.go file
- In terminal `go mod init hello`
- Then the code in main.go file.
- To run the file:
	- In terminal, `go run fileName.go` in our case, `go run main.go`
	- So this just compiles and runs the program.
	- We are not going to get any executables.
-  To view the commands -  `go help`
- To view more information about a particular command:
	- `go help commandName` in our case let's find out more information about, `go help run'
- To view the current GOPATH:
	- `go env GOPATH` - Case sensitive command.
	- cd to the GOPATH
	- Do `ls`
	- cd pkg, then ls
	- cf mod, then ls
	- We can find few files like cache, github.com, gopkg.in and few more.
	- These are the default packages.

### 02 Variables

- Created main.go file
- In terminal `go mod init variables`
- Then the code in main.go file.

* Resources
	- [Documentation of fmt package](https://pkg.go.dev/fmt)
	- [Types in Golang](https://pkg.go.dev/builtin#pkg-types)
	- [Documentation Lexical elements](https://golang.org/ref/spec#Lexical_elements)

### 03 User Input

- Created main.go file
- In terminal `go mod init userinput`
- Then the code in main.go file.

* Resources
	* To search any package in Go: [https://pkg.go.dev/](https://pkg.go.dev/)
	* [bufio](https://pkg.go.dev/bufio) package to get the input from user and other source of inputs.
	* [os](https://pkg.go.dev/os) is also one of the useful packages.

### 04 Conversion

- Created main.go file
- In terminal `go mod init conversion`
- Then the code in main.go file.

* Resources
	* [strconv](https://pkg.go.dev/strconv) package is used to convert string to various types.
	* [strings](https://pkg.go.dev/strings) package is very powerful package in Golang.