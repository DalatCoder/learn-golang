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

### 1.2. Short declaration operator

Look at `go spec`

- Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

`identifier = letter { letter | unicode_digit } . `

Read more about `ebnf` for this syntax.

Keyword: the word that are reversed and may not be used as identifiers.
Operator and punctuation

Short declaration opertor

```go
package main

import "fmt"

func main() {
  x := 42
  fmt.Println(x)

  x = 99
  fmt.Println(x)

  y := 100 + 24
  fmt.Println(y)
}
```

It declares a variable `x` and assign value of `42` to it.
When we want to assign a new value, we just use the assign operator for this `=`.

Summary

- keywords:
- operator:
- operand:
- statement:
  In programming, a statement is the smallest standalone element of a program that expresses some action to be carried out. It is an
  instruction that commands the computer to perform a specificed action. A program is formed by a sequence of one or more statements.

- expression: in programming an expression is a combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and produce another value. For example, 2 + 3 is an expression which evalute to 5.

### 1.3. Exploring types

- Declare a variable and assign a value at the same time: `x := 42` (short declaration operator)
- Declare variable with `var`

```go
var y = 43
fmt.Println(y)
```

When we want to declare a variable outside of a function body, we use `var`. Otherwise, we use short declaration syntax.

And, declare there is a variable with the identifier `z` and that the variable with the identitier z is of type int, assign the zero value
of type int to `z`

`var z int`

Look at the `golang` spec for `zero` value.

- boolean: false
- integer: 0
- float: 0.0
- string: ""
- pointers, functions, interfaces, slices, channels, map: nil

### 1.4. Exploring types (continue)

- Declare a variable is of a certain type it can only hold values of that type

- `primitive data types`

  - In computer science, primitive data type is either of the following:

    - `a basic type` is a data type provided by a programming language as a basic building block. Most languages allow more complicated `composite` types to be constructed starting from basic types.
    - `a built in type` is a data type for which the programming language language provides built-in support.

  - In most programming languages, all basic data types are built-in. In addition, many languages also provide a set of composite data types. Opinions vary as to whether a built-in type that is not basic should be considered "primitive".

- `composite data types`
  - In computer science, a `composite data type` or `compound data type` is any data type which can be constructed in a program using the programming lanuage's primitive data types and other composite types. It is sometimes called a `structure` or `aggregate data type`.

```go
package main

import "fmt"

var y = 42
var z = "tronghieu"
var a = `"tronghieu"`

func main() {
  fmt.Printf("%T\n", y)
  fmt.Printf("%T\n", z)
}
```

### 1.5. Zero value

Understand `zero value`

- `false` for booleans
- 0 for integers
- 0.0 for floats
- "" for strings
- `nil` for

  - pointers
  - functions
  - interfaces
  - slices
  - channels
  - maps

Use short declaration operator as much as possible
Use `var` for

- `zero` value
- `package scope`

### 1.6. Create your own types

Go is all about types

```go
package main

import (
  "fmt"
)

var a int
type hotdog int

var b hotdog

func main() {
  a = 42

  fmt.Println(a)
  fmt.Printf("%T\n", a)

  b = 42
  fmt.Println(b)
  fmt.Printf("%T\n", b) // main.hotdog

  a = b // wrong
}
```

We create our own type called `hotdog` and the underlying type is an `int`

### 1.7. Conversion, not casting

```go
a = int(b)
```

## 2. Exercises number #1

### 2.1. Exercises #1

Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z"

- 42
- "James Bond"
- true

Now print the values stored in those variables using

- a single print statement
- multiple print statement

```go
package main

import "fmt"

func main() {
  x := 42
  y := "James Bond"
  z := true

  fmt.Println(x, y, z)
}
```

### 2.2. Exercises #2

Use var to DECLARE three variables. The variables should have package level scope.
Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
variables and make sure the variables store values of following TYPE

- identfier "x" type int
- identifier "y" type string
- identifier "z" type bool

In func main

- Print out the values for each identifier
- The compiler assigned values to the variables. What are these values called?

```go
package main

import (
  "fmt"
)

var x int
var y string
var z bool

func main() {
  // zero values
  fmt.Println(x) // 0
  fmt.Println(y) // ""
  fmt.Println(z) // false
}
```

### 2.3. Exercises #3

Using the code from the previous exercise,

- at the package level scope, assign the following values to the three variables

  - for x assign 42
  - for y assign "James Bond"
  - for z assign true

- in func main
  - use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
    returned value of TYPE string using the short declaration operator to a VARIABLE
    with the IDENTIFIER "s"

```go
package main

import (
  "fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
  s := fmt.Sprintf("%v %v %v", x, y, z)
  fmt.Println(s)
}
```

### 2.4. Exercises #4

- Create your own type. Have the underlying type be an `int`
- Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR" keyword
- in func main

  - print out the value of the variable x
  - print out the type of the variable x
  - assign 42 to the VARABLE x using the = OPERATOR
  - print out the value of variable x

```go
package main

import (
  "fmt"
)

type nth int

var x nth

func main() {
  fmt.Println(x)
  fmt.Printf("%T\n", x)

  x = 42
  fmt.Println(x)
}
```

## 3. Programming Fundamentals

See more at Go language specification

### 3.1. Boolean type

```go
package main

import "fmt"

var x bool

func main() {
  fmt.Println(x) // false

  x = true
  fmt.Println(x) // true

  a := 7
  b := 42
  fmt.Println(a == b)
}
```

### 3.2. How computers work

- computers run on electricity
- electricity can be ON or OFF
- coding schemes
  - assign meaning to ON or OFF state

ON & OFF, 1 & 0, Binary Digits, Bits and Machine Language are all words used to refer to this
idea that, within a computer , it's all nothing but a bunch of ZERO's and ONE's, or switches that are ON or OFF,
it's all just a bunch of Binary Digits, or Bits, that's the language which computer speak,
it's a machine language.

Circuits, Switches, Transitors, and even "gates" are all words used to refer to this thing
within a computer that can either be ON or OFF. It's a circuit, it's a switch, it's a gate
that can either be OPENED or CLOSED, it's a transitor - you will learn that people use of all those
words to talk about this same thing, this ability of computers to store ON/OFF states.

### 3.3. Numeric types

```go
package main

import "fmt"

func main() {
  x := 42
  y := 42.45678

  fmt.Println(x)
  fmt.Printf("%T\n", x) // int
  fmt.Println(y)
  fmt.Printf("%T\n", y) // float64
}
```

Summary

- integers:

  - numbers without decimals (whole number)

- int & uint

  - implementation-specific sizes

- all numeric types are distinct except

  - `byte` which is an alias for `uint8`
  - `rune` which is an alias for `int32` (1 character in UTF8 need from 1 to 4 byte)

- types are unique
  - this is static programming language
  - conversions are required when different numeric types are mixed in an
    expression or assignment. For instance, `int32` and `int` are not the same type even though they
    may have the same size on a particular architecture.
- rule of thumb: just use `int`
- floating point: `float64`

### 3.4. String types

```go
package main

import "fmt"

func main() {
  s := "Hello, World"

  fmt.Println(s)
  fmt.Printf("%T\n", s)

  s1 := `
    Hello
    World
  `
  fmt.Println(s1)
  fmt.Printf("%T\n", s1)
}
```

A string value is a (possibly empty) sequence of bytes. String are immutable.

```go
package main

import "fmt"

func main() {
  s := "Hello World"

  bs := []byte(s)

  fmt.Println(bs)
  fmt.Printf("%T\n", bs) // []uint8

  for i := 0; i < len(s); i++ {
    fmt.Printf("%#U\t", s[i])
  }

  for i, v := range s {
    fmt.Println(i, v)
  }
}
```

### 3.5. Constants

- A simple, unchanging value
- Only exist at compile time
- there are TYPED and UNTYPED constants

  - `const hello = "Hello World"`
  - `const typedHello string = "Hello Wolrd"`

- Untyped constant
  - a constant that does not have yet have a fixed type
    - `constant of a kind`
    - not yet forced to obey the strict rules that prevent combining differently typed values
  - an untyped constant can be implicitily converted by the compiler

```go
package main

import "fmt"

const a = 42
const b = 42.78
const c = "James Bond"

const (
  d = 3
  e = 3.14
  f = "Hello"
)

func main() {
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Printf("%T\n", a) // int
  fmt.Printf("%T\n", b) // float64
  fmt.Printf("%T\n", c) // string

  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)
}
```
