# Go: The Complete Developer's Guide

## 1. Hello World Application

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

### 1.1. How do we run the code in the project?

- `go run main.go`
- `go build main.go && ./main.exe`
- `go run *.go`
- `go test ./...`

Go CLI

- `go build`: Compiles a bunch of `go` source code files
- `go run`: Compiles and executes
- `go fmt`: Formats all the code in each file in the current directory
- `go install`: Compiles and "installs" a package
- `go get`: Downloads the raw source code of someone else's package
- `go test`: Runs any tests associated with the current project

### 1.2. What does `package main` mean?

`package == project == workspace`

A package can have multiple related files inside of it.

Type of packages

- `executable`: Generates a file that we can run: `package main`
- `reusable`: Code used as `helpers`. Good place to put reusable logic

### 1.3. What does `import "fmt"` mean?

Provide access from our package to another packages. `golang.org/pkg`

### 1.4. What's that `func` thing?

Function.

### 1.5. How is the `main.go` file organized?

- package declaration: `package main`
- import other packages that we need: `import fmt`
- declare functions, tell Go to do things: `func main() {}`