# Golang Udemy

## Environment

Go workspace

- one folder - any name, any location

  - bin
  - pkg
  - src

          - github.com
              - <github.com username>
                  - folder with code for project / repo
                  - folder with code for project / repo
                  - folder with code for project / repo

- namespacing
- `go get`

  - package management

- `GOPATH`

  - points to your go workspace

- `GOROOT`

  - points to your binary installation of Go

Go commands

- `go version`
- `go env`
- `go help`
- `go fmt`
- `go run`
  - needs ag file name, eg, `go run main.go`
  - `go run <filename>`
  - go run \*.go
- `go build`

  - for an executable:
    - builds the file
    - reports errors, if any
    - throws away binary
  - for a package:
    - builds the file
    - reports errors, if any
    - throws away binary

- `go install`
  - for an executable:
    - compiles the porgram (builds it)
    - names the executable the folder name holding the code
    - puts the executable in `wordspace/bin`
      - `$GOPATH/bin`
  - for a package:
    - compiles the package (builds it)
    - puts the execuatable in `workspace / pkg`
      - `$GOPATH/pkg`
    - makes it an archive file

Package management

Starting from version `1.13.0`, go module

Creating a go module

- `go mod init` create a new module, initializing the `go.mod` file that describes it.
- `go build`, `go test` and other package-building commands add new depencecies to `go.mod`
- `go list -m all` prints the current module's dependencies
- `go get` changes the required version of a dependency (or adds a new dependency)
- `go mod tidy` removes unused dependencies

Example

- Create folder `01.happydog`
- Create 2 files

  - `hello.go`
  - `hello_test.go`

`go test`

Let's make the current directory the root of a module by using `go mod init` and then try `go test` again

```sh
    go mod init example.com/username/repo
```

And now we can run `go test`

Adding a dependency

`import "rsc.io/quote"`

Go to `godoc.org`

- `go mod init example/username/repo`
- `cat go.mod`

- `go list -m all`

Both `go.mod` and `go.sum` should be checked into version control

Upgrade dependencies

## 1. Variables, Values and Types

Every program must to have package `main`

```go
package main

func main()  {

}
```

`func main` is the entry point of the program

Control flow

- (1) sequence
- (2) loop: iterative
- (3) conditional

```go
package main

import "fmt"

func main()  {
  fmt.Println("Hello World")

  foo()

  for i := 0; i < 100; i++ {
    if i % 2 == 0 {
      fmt.Println(i)
    }
  }

}

func foo() {
  fmt.Println("I'm in foo")
}
```

basic program structure

- package `main`
- func `main`

  - entrypoint to your program
  - when your code exits func `main`, your program is over

### 1.1. Introduction to packages

talk about packages and modularize in your code.

We organize our data or information n a computer with folders. In a program, we use packages

`https://godoc.org/fmt`

Look at the documentation of `fmt` packages

Signature of one method:

- `func Println(a ...interface{}) (n int, err error)`

Here we have an empty `inteface`, which means they'll take a value of any type and an unlimitted number of them

```go
fmt.Println("Hello world", 42, true);
```

Catch up returns values

```go
n, err := fmt.Println("Hello world", 42, true);
fmt.Println(n) // number of bytes
fmt.Println(err) // nil
```

Throw away the return by using underscore `_`

```go
n, _ := fmt.Println("Hello world", 42, true);
fmt.Println(n) // number of bytes
```

Summary

- variadic parameters:
  - the `...<some type>` is how we signify a `variadic parameter`
  - the type `interface{}` is the empty `interface`
    - everything is of type `interface{}`
  - so the parameter `...interface{}` means `give me as many arguments of any type as you'd like`
- throwing away returns
  - use the `_` character to throw away returns
- you can't have unused variables in your code
  - this is code pollution
  - the compiler doesn't allow it
- we use this notation in `go`
  - `package`.`Identifier`
    - ex: `fmt.Println()`
      - we would read that: "From package `fmt` I am using the `Println` func"
- An identifier is a name of the variale, constant, func
- packages
  - code that is already written which you can use
  - `import`
