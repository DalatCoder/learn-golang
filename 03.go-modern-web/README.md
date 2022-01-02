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

### 1.4. Receivers Structs with Functions

Structs can have functions associated with them.

```go
package main

type myStruct struct {
    FirstName string
}

func main() {
    m := myStruct { "Hieu" }
    n := myStruct { "Ha" }
}
```

Create a new function and attach it to the struct through the `receiver`.

`(m *myStruct)` is called a `receiver` and it ties this function to the `type of myStruct` because
I'm using a `pointer`. This `pointer` of `*myStruct` point to the `myStruct` struct. That mean I
can access to the information from `myStruct` through this `receiver`.

```go
func (m *myStruct) printFirstName() string {
    return m.FirstName
}
```

Using `receiver` with `pointer` to `save` the change.

```go
package main

import (
  "fmt"
)

type Person struct {
  FirstName string
  LastName  string
}

func (p Person) changeFirstName() {
  p.FirstName = "Hieu Normal"
}

func (p *Person) changeFirstNamePointer() {
  (*p).FirstName = "Hieu Pointer"
}

func main() {
  p := Person{
    FirstName: "Ha",
    LastName:  "Nguyen",
  }

  fmt.Println("Init", p.FirstName)

  p.changeFirstName()
  fmt.Println("Change first name normal", p.FirstName)

  p.changeFirstNamePointer()
  fmt.Println("Change first name pointer", p.FirstName)
}

```

### 1.5. Map & Slice

Half of what you doing when you're writing programs is figuring out what kind of data structure you want
to store your data in.

- simple variables
- structs

Map data structure (`not to be sort`)

Map is `immutable`, never bother to passing a pointer to a `map` to change it. You can just pass
that `map` itself and that map will remain constant no matter where in the program it is accessed.

```go
package main

func main() {
    myMap := make(map[string]string) // best practice
    myMap := make(map[string]interface{}) // put any thing to map (not recommend)

    myMap["dog"] = "kiki"
    myMap["other-dog"] = "hihi"

    log.Println(myMap["dog"], myMap["other-dog"])
}
```

Slice data struct (`can be sort`)

Concept of `array`. But in `go`, we almost never use `array`, we use `slice`.

```go
package main

func main() {
    mySlice := []string // a slice of string

    mySlice = append(mySlice, 2)
    mySlice = append(mySlice, 1)
    mySlice = append(mySlice, 3)

    sort.Ints(mySlice)
}
```

Half of what you doing when you're writing programs is figuring out what kind of data structure you want
to store your data in.

- simple variables
- structs

Map data structure (`not to be sort`)

Map is `immutable`, never bother to passing a pointer to a `map` to change it. You can just pass
that `map` itself and that map will remain constant no matter where in the program it is accessed.

```go
package main

func main() {
    myMap := make(map[string]string) // best practice
    myMap := make(map[string]interface{}) // put any thing to map (not recommend)

    myMap["dog"] = "kiki"
    myMap["other-dog"] = "hihi"

    log.Println(myMap["dog"], myMap["other-dog"])
}
```

Slice data struct (`can be sort`)

Concept of `array`. But in `go`, we almost never use `array`, we use `slice`.

```go
package main

func main() {
    mySlice := []string // a slice of string

    mySlice = append(mySlice, 2)
    mySlice = append(mySlice, 1)
    mySlice = append(mySlice, 3)

    sort.Ints(mySlice)
}
```

### 1.6. Decision structures

- If statement
- Switch statement

### 1.7. Loop and ranging over data

- For loop
- Range over data (loop over every entry of a slice)

### 1.8. Interfaces

In really simplest terms, an interface is nothing more than a kind of contract.

An `interface` says, OK, any type that uses this interface has to follow these rules.

Create more `generic` function that accept a type, an interface type, rather than a
specific type like `Dog` or `Gorilla`.

The key thing to remember is that once you define an `interface`, anything else can implement that interface,
just by implementing the required methods.

```go
package main

type Animal interface {
    Says() string
    NumberOfLegs() int
}

type Dog struct {
    Name string
    Breed string
}

type Gorilla struct {
    Name string
    Color string
    NumberOfTeeth int
}

func main() {
    dog := Dog {
        Name: "kiki",
        Breed: "kikikaka"
    }

    PrintInfo(dog)

    gorilla := Gorilla {
        Name: "King Kong",
        Color: "black",
        NumberOfTeeth: 32
    }

    PrintInfo(gorilla) // wrong
}

func (d Dog) Says() string {
    return "Woof"
}

func (d Dog) NumberOfLegs() int {
    return 4
}

func PrintInfo(a Animal) {
    log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs, "legs")
}
```
