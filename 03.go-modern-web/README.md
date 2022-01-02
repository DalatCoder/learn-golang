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

### 1.3. Type & Struct

Sometimes the primitive variables are not enough to do what you want them to do.

So, example of store a person in database, we could define those variales

```go
package main

import (
    "log"
    "time"
)

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time
```

We define new `type`.

We define struct with `Capitalize` because we want to use that struct in another package.

- `camelCase`: private, only used inside the current package
- `CapitalCase`: public, can used outside the current package

```go
package main

import (
    "log"
    "time"
)

type User struct {
    FirstName string
    LastName string
    PhoneNumber string
    Age int
    BirthDate time.Time
}

func main() {
    user := User {
        FirstName: "Hieu",
        LastName: "Nguyen Trong",
        PhoneNumber: "0123456789",
        Age: 22
    }

    log.Println(user.FirstName)
}
```
