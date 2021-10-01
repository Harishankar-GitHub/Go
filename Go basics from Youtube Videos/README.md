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