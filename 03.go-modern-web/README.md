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

### 1.9. Packages

Define our own package.

So packages are incredibly useful, incredibly powerful, a nice way of getting our code organized into
logical groupings.

Convention to name the package like this

- `go mod init github.com/DalatCoder/myniceprogram`
- create new directory: `helpers`
  - create new file: `package helpers`

### 1.10. Channels

Channel are unique to Go.

They are a means of sending information from one part of your program to another part of your program, very easily.

- Passing information from a function to another function through arguments
- Passing information form a function to another function using pointers
- Creating a channel

`intChan := make(chan int) // make a channel of int`. And all that says is I'm creating a channel,
a place to send information which will be received in one or more places in my program, and that channel can only
hold `int`.

```go
package main

const numPool = 1000

func RandomNumber(n int) int {
    rand.Seed(time.Now().UnixNano())
    value := math.Intn(n)
    return value
}

func CalculateValue(intChan chan int) {
    randomNumber := RandomNumber(numPool)
    intChan <- randomNumber // put random number to our channel
}

func main() {
    intChan := make(chan int) // make a channel of int
    defer close(intChange)

    // Start a goroutine a run the function CalculateValue in this go roroutine
    go CalculateValue(intChan)

    // Get value from the channel (listen for changing event)

    num := <- intChan // Assign whatever value come from channel to num variable
    log.Println(num)
}
```

### 1.11. Readig and Writing JSON

FirstName string `json:"first_name"`, define a `field name` and the matching name in `response json`

```go
package main

type Person struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
}

func main() {
    myJson := `
        [
            {
                "first_name": "Hieu",
                "last_name": "Nguyen"
            },
            {
                "first_name": "Ha",
                "last_name": "Nguyen"
            }
        ]
    `

    var unmarshalled []Person

    err := json.Unmarshall([]byte(myJson), &unmarshalled)
    if err != nil {
        log.Println("Error marshalling json", err)
    }

    log.Printf("unmarshalled: %v", unmarshalled)


    // Write json from a struct
    mySlice := []Person{
        { "Hieu", "Nguyen" },
        { "Ha", "Nguyen" }
    }

    // with format
    newJson, err := json.MarshalIndent(mySlice, "", "  ")

    // no format, use in production env
    newJson, err := json.Marshal()

    if err != nil {
        log.Println("Error marshalling json", err)
    }

    fmt.Println(string(newJson))
}
```

### 1.12. Writing tests

When we write code, we want to write tests to ensure that our code actually does what it's supposed to do.

In file `main.go`

```go
package main

func main() {
    result, err := divide(100.0, 10.0)
    if err != nil {
        log.Println(err)
        return
    }

    log.Println(result)
}

func divide(x, y float64) (float32, error) {
    var result float32

    if y == 0 {
        return result, errors.New("cannot divide by 0")
    }

    result = x / y
    return result, nil
}
```

In file `main_test.go`

2 ways to write tests

- one way slow
- one way efficient

The first way: write every test function for every case

```go
package main

func TestDivide(t *testing.T) {
    _, error := divide(10.0, 1.0)

    if err != nil {
        t.Error("Got an error when we should not have")
    }
}

func TestBadDivide(t *testing.T) {
    _, error := divide(10.0, 0)

    if err == nil {
        t.Error("Did not get an error when we should have")
    }
}
```

The second way: write a table test

```go
package main

var tests = []struct {
    name string
    dividend float32
    divisor float32
    expected float32
    isErr bool
}{
    { "valid-data", 100.0, 10.0, 10.0, false },
    { "invalid-data", 100.0, 0.0, 0, true },
    { "expect-5", 100.0, 20.0, 5.0, false },
    { "expect-fraction", -1.0, -777.0, 0.0012870013, false },
}

func TestDivision(t *testing.T) {
    for _, tt := range tests {
        got, err := divide(tt.dividend, tt.divisor)

        if tt.isErr {
            if err == nil {
                t.Error("expected an error bud did not get one")
            }
        } else {
            if err != nil {
                t.Error("did not expect an error but got one", err.Error())
            }
        }

        if got != tt.expected {
            t.Errorf("expected %f but got %f", tt.expected, got)
        }
    }
}
```

Run test: `go test -v`
