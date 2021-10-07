## Basics of Golang

> **IDE - Visual Studio Code**

1. **Hello**
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

2. **Variables**
- Created main.go file
- In terminal `go mod init variables`
- Then the code in main.go file.

* Resource
	- [Documentation of fmt package](https://pkg.go.dev/fmt)
	- [Types in Golang](https://pkg.go.dev/builtin#pkg-types)
	- [Documentation Lexical elements](https://golang.org/ref/spec#Lexical_elements)

3. **User Input**
- Created main.go file
- In terminal `go mod init userinput`
- Then the code in main.go file.

* Resource
	* To search any package in Go: [https://pkg.go.dev/](https://pkg.go.dev/)
	* [bufio](https://pkg.go.dev/bufio) package to get the input from user and other source of inputs.
	* [os](https://pkg.go.dev/os) is also one of the useful packages.

4. **Conversion**
- Created main.go file
- In terminal `go mod init conversion`
- Then the code in main.go file.

* Resource
	* [strconv](https://pkg.go.dev/strconv) package is used to convert string to various types.
	* [strings](https://pkg.go.dev/strings) package is very powerful package in Golang.

5. **Time**
- Created main.go file
- In terminal `go mod init mytime`
- Then the code in main.go file.

* `go env` - Env prints Go environment information.
* `go help env` - More information about `go env` command.

- In `go env` command, there are a lot of information about the Go environment.
- One such is **GOOS=windows** (Because I use a Windows machine)
- This will be different depending on the Operating System.
- In my case, I'm in a Windows machine.
- I can generate the build package for any Operating System from Windows.
- Command to do that: (Doesn't work for me. Need to check)
	- `GOOS="windows" go build` - For Windows
	- `GOOS="linux" go build` - For Linux
	- `GOOS="darwin" go build` - For Mac (darwin is considered as Mac)

* Resource
	* [time](https://pkg.go.dev/time) package in Golang.

6. **Pointers**
- Created main.go file
- In terminal `go mod init pointers`
- Then the code in main.go file.

>- Generally in programming languages, when a variable or constant is created, the variable has the reference to the address.
>- When we pass these variables to methods/functions etc., sometimes a copy of value is passed instead of the actual value.
>- This is where the ***Pointer concept in Golang*** plays a major role.
>- Pointer provides 100% guarantee that the actual value is passed on every time.

7. **Arrays**
- Created main.go file
- In terminal `go mod init arrays`
- Then the code in main.go file.

>- In other programming languages, Arrays are one of the most used data structures.
>- But in Golang, the language specification holds our hand a lot and don't allow us  to use too much of the array.
>- But under the hood, there are other datatypes which actually utilises array but gives us lot more freedom than that.

* Resource
	* [Array Types](https://golang.org/ref/spec#Array_types) in Golang.

8. **Slices**
- Created main.go file
- In terminal `go mod init slices`
- Then the code in main.go file.

>- Slice is the ***mostly used datatype*** in Golang.
>- Slices are actually under the hood arrays.
>- These arrays once they get an abstraction layer and some more features, they are called as Slices.
>- In Golang, Slices are used more as they are more powerful.

* Resource
	* [Slice Types](https://golang.org/ref/spec#Slice_types) in Golang.

9. **Maps**
- Created main.go file
- In terminal `go mod init maps`
- Then the code in main.go file.

* Resource
	* [Map Types](https://golang.org/ref/spec#Map_types) in Golang.

10. **Structs**
- Created main.go file
- In terminal `go mod init structs`
- Then the code in main.go file.

> ***Note:***
	>- ***No Inheritance in Golang.***
	>- ***No parent or super concepts in Golang.***

* Resource
	* [Struct Types](https://golang.org/ref/spec#Struct_types) in Golang.

11. **If-Else**
- Created main.go file
- In terminal `go mod init ifelse`
- Then the code in main.go file.

> In If-Else syntax, the curly braces must be on the same line.
> Correct Syntax:
> ```
> if condition {
>	//stmt
> }
> ```
> Incorrect Syntax:
> ```
> if condition
> {
> //stmt
> }
> ```

* Resource
	* [If-Else](https://golang.org/ref/spec#If_statements) in Golang.

12. **Switch case**
- Created main.go file
- In terminal `go mod init switchcase`
- Then the code in main.go file.

* Resource
	* [Switch case](https://golang.org/ref/spec#Switch_statements) in Golang.

13. **Loops**
- Created main.go file
- In terminal `go mod init loops`
- Then the code in main.go file.

14. **Functions**
- Created main.go file
- In terminal `go mod init functions`
- Then the code in main.go file.

* Resource
	* [Functions](https://golang.org/ref/spec#Function_types) in Golang.
	* [Function Declarations](https://golang.org/ref/spec#Function_declarations) in Golang.

15. **Methods**
- Created main.go file
- In terminal `go mod init methods`
- Then the code in main.go file.

* Resource
	* [Method Declarations](https://golang.org/ref/spec#Method_declarations) in Golang.

16. **Defer**
- Created main.go file
- In terminal `go mod init defer`
- Then the code in main.go file.

* Resource
	* [Defer Statements](https://golang.org/ref/spec#Defer_statements) in Golang.