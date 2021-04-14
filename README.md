# Go

- [Go Homepage](https://golang.org/)
- [Package Docs](https://golang.org/pkg/)
- [Other Docs](https://golang.org/doc/)

- Useful website for Go: [learn.go.dev](https://learn.go.dev/)

+ Installation - [Click here](https://golang.org/doc/install) or [Click here](https://golang.org/dl/)

After installing, check the version in cmd/Power Shell - `go version`

- To view the commands - `go help`
---

### Basics of Go Language

> **IDE - Visual Studio Code**

1. [Hello World](https://github.com/Harishankar-GitHub/Go/blob/main/1%20Hello%20World/helloworld.go)
2. [Datatypes](https://github.com/Harishankar-GitHub/Go/tree/main/2%20Datatypes)
3. [Conrol Structures](https://github.com/Harishankar-GitHub/Go/tree/main/3%20Control%20Structures)
4. [Loops](https://github.com/Harishankar-GitHub/Go/tree/main/4%20Loops)
5. [Errors](https://github.com/Harishankar-GitHub/Go/blob/main/5%20Errors/errors.go)
6. [Pointers](https://github.com/Harishankar-GitHub/Go/blob/main/6%20Pointers/pointers.go)
7. [Structures](https://github.com/Harishankar-GitHub/Go/blob/main/7%20Structures/structures.go)
8. [Slices and Maps](https://github.com/Harishankar-GitHub/Go/tree/main/8%20Slices%20and%20Maps)
9. [Methods](https://github.com/Harishankar-GitHub/Go/blob/main/9%20Methods/methods.go)
10. [Interfaces](https://github.com/Harishankar-GitHub/Go/tree/main/10%20Interfaces)

### Organizing the code

- Go uses 2 environment variables to manage the projects.
	> `GOPATH` and `GOROOT`

- The projects must exist within the project directory structure to use external dependencies.

### Tools for Go Development

- Go provides several tools that help working the your projects.
- There are tools for ***linting, vetting, and formatting*** the code.

### Vendoring

- Vendoring means that the library code that the project depends on will be checked in with the project.
- [Glide](https://glide.readthedocs.io/en/latest/) is a tool to manage the vendor dependencies.

### Testing

- Tests in Go are simple to write and are included directly in your codebase next to other source code files.
- There are ***no asserts***, only failures.

---