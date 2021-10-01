# Go

- [Go Homepage](https://golang.org/)
- [Package Docs](https://golang.org/pkg/)
- [Other Docs](https://golang.org/doc/)

- Useful website for Go: [learn.go.dev](https://learn.go.dev/)
	- [Go by Example](https://gobyexample.com/)
	- [Web Dev](https://gowebexamples.com/)

+ Installation - [Click here](https://golang.org/doc/install) or [Click here](https://golang.org/dl/)

> After installing, check the version in cmd/Power Shell - `go version`

> To view the commands - `go help`
---

## Resources

- [Youtube playlist for Golang](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)

## Golang
- Golang is Compiled
	- Go tool can run file directly, without VM
	- Executables are different from OS
	- Executables are exported at the compile time.
	- Go gives the options to produce the compiled output or the final executables for windows/linux/mac.
- What is Golang ? Where can it be used ?
	- System Apps to Web Apps - Cloud
- Object Oriented ?
	- Yes and No
	- According to OOPS, a language should have Classes, Objects etc.
	- Golang doesn't have those. Instead it has Structs, which is an altenative.
	- Whatever one sees on the screen, the code just does that.
	- This is how Golang is designed.
- Missing
	- Few features that are available in other languages are missing in Golang.
	- One such thing is try - catch block.
	- And it is not needed in Go.
	- No need to write semi colon while coding in Golang.
		- But the document says semi colon is required.
		- Even if we put semicolons at the end of the statements (which is still correct), as soon as we save the file, the [Go Extension](https://code.visualstudio.com/docs/languages/go) that is in the Visual Studio Code removes it.
		- Lexer does a lot of work. [Documentation](https://golang.org/ref/spec#Semicolons)
- Types
	- Case insensitive, almost
	- We can also differenciate a method, variable etc., if it is private or public.
	- If the first letter is capitalized, then it is Public.
	- Variable type should be known in advance.
		- But there are some syntax which allows us to declare a variable without specifying the type and it will predict the type on the go.
	- Everything in the Golang is a Type.

	* Few types
		- String
		- Bool
		- Integer
			- Aliases are involved too - `uint8`, `unit64`, `int8`, `int64`, `uintptr`
		- Floating
			- `float32`, `float64`
		- Complex
		- Arrays
		- Slices
		- Maps
		- Structs
		- Pointers

	> Almost everything, Functions, CHannels etc., are Types in Golang.