# Building Modern Web Applications in Go

Why Go?

- Compiles to a single binary file
- No runtimes to worry about (because of the binary file)
- Statically typed, so no surprises at run time
- Object oriented (sort of)
- Concurrency
- Cross platform
- Excellent package management & testing built-in

## 1. Overview of the Go language

### 1.1. Variables & Function

File structure

- end with `.go`
- begin of file with `package declaration`: any name, but the convention to call `main` for the main application.
- must be atleast 1 function, called `main` (also the `entry point` of the application)

Create variable:

- `var camelCase <type>`
- short declaration `x := expression`

`log` package

- `import "log"`
- `log.Println()`

Variable types

- `var i int`: zero value: 0
- `var s string`: zero value: ""

Go have first-class functions, a functions can return multiple values

- `func saySomething(s string) (string, string) { return "Hieu", "Ha" }`
- `s, ss := saySomething("")`

### 1.2. Pointers

Incredibly useful.

A pointer points to a specific location in memory and gives you a mean of getting  
that particular location in memory.

```go
package main

import "log"

func main() {
    s := "green"
    log.Println("My string is set to", s)

    changeUsingPointer(&s)
    log.Println("My string is set to", s)
}

func changeUsingPointer(s *string) {
    *s := "red"
}
```
