# Golang Udemy

- go doc
- go effective
- go spec
- go playground
- common pattern: `keyword identifier type`
  - `var x int`
  - `type person struct{}`
  - `type human interface{}`

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

### 3.6. Iota

Predeclare identifier

Within a constant declaration, the predeclare identifier `iota` represents successive
untyped integer constants. It is reset to 0 whenever the reserved word `const` appears
in the source and increments after each `ConstSpect`. It can be used to construct a
set of related constants.

If you need something automatically increment by one.

```go
package main

import "fmt"

const (
  a = iota
  b = iota
  c = iota
)

const (
  d = iota
  e
  f
)

func main() {
  fmt.Println(a) // 0
  fmt.Println(b) // 1
  fmt.Println(c) // 2
  fmt.Printf("%T\n", a) // int
  fmt.Printf("%T\n", b) // int
  fmt.Printf("%T\n", c) // int

  // Auto reset
  fmt.Println(d) // 0
  fmt.Println(e) // 1
  fmt.Println(f) // 2
}
```

### 3.7. Bit Shifting

Two operator

- `<<`
- `>>`

Example

```go
package main

import (
  "fmt"
)

func main() {
  // 100 = 4
  x := 4
  fmt.Printf("%d\t\t%b", x, x) // Decimal and binary

  // 1000 = 8
  y := x << 1
  fmt.Printf("%d\t\t%b", y, y) // Decimal and binary
}
```

Apply `iota` the creative way

```go
package main

import (
  "fmt"
)

func main() {
  kb := 1024
  mb := 1024 * kb
  gb := 1024 * mb

  fmt.Printf("%d\t\t%b", kb, kb)
  fmt.Printf("%d\t\t%b", mb, mb)
  fmt.Printf("%d\t\t%b", gb, gb)
}
```

Using `iota` for automatically calculate

```go
package main

import (
  "fmt"
)

const (
  _ = iota
  kb = 1 << (iota * 10)
  mb = 1 << (iota * 10)
  gb = 1 << (iota * 10)
)

func main() {
  fmt.Printf("%d\t\t%b", kb, kb)
  fmt.Printf("%d\t\t%b", mb, mb)
  fmt.Printf("%d\t\t%b", gb, gb)
}
```

## 4. Control Flow

### 4.1. for init; condition; post

There is no `while` loop in `Go`

`for` loop: `for init; condition; post`

```go
package main

import (
  "fmt"
)

func main() {

  // for init; condition; post {}
  for i := 0; i <= 10; i++ {
    fmt.Println(i)
  }
}
```

### 4.2. Nesting loop

```go
package main

import (
  "fmt"
)

func main() {

  for i := 0; i <= 10; i++ {

    for j := 0; j <= 5; j++ {
      fmt.Println(i * j)
    }
  }
}
```

### 4.3. For statement

A `for` statement specifies repeated execution of a block. There are three forms

- the iteration may be controlled by a single condition
- a "for" clause
- a "range" clause

For statement with single condition

```go
for a < b {
  a *= 2
}
```

For statement with `for` clause

```go
for i := 1; i < 10; i++ {
  fmt.Println(i)
}
```

Infinite loop statement

```go
for {
  // Listen for incomming request on TCP
}
```

### 4.4. If statement

- bool : true / false
- the not operator: !
- initialization statement
- if / else
- if / else if / else

```go
package main

import (
  "fmt"
)

func main() {
  if true {
    fmt.Println("001")
  }

  if false {
    fmt.Println("001")
  }

  if !true {
    fmt.Println("001")
  }

  if !false {
    fmt.Println("001")
  }

  if x := 2; x % 2 == 0 {

  }

  fmt.Println(x) // error: x is not avaiable

  if true {}
  else if false {}
  else {}
}
```

### 4.5. Switch statement

```go
package main

import (
  "fmt"
)

func main() {
  variable := true

  switch variable {
    case false:
      fallthrough
    case true:
      fallthrough
    case "a", "b", "c":
      fallthrough
    default case  :
  }
}
```

### 4.6. Conditional logic operator (&& || !)

## 5. Grouping data

## 5.1. Array

Arrays are useful when planning the detailed layout of memory and sometimes can help
avoid allocation, but **primiarily they are a building block for slices**.
**Use slices instead**

```go
package main

import (
  "fmt"
)

func main() {
  var x [3]int

  fmt.Println(x) // [0, 0, 0]

  x[2] = 1
  fmt.Println(x) // [0, 0, 1]
}
```

## 5.2. Slice

Composite literal: `x := type{values}`.

A SLICE allows you to group together VALUES of the same TYPE

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3, 4, 5}

  fmt.Println(x) // [1 2 3 4 5]
}
```

### 5.2.1 Slice - for range

If you're looping over an array, slice, string, or map, or reading from a channel, a `range` clause can manage
the loop:

```go
  for key, value := range oldMap {
    newMap[key] = value
  }
```

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3, 4, 5}
  fmt.Println(len(x)) // 5

  for i, v := range x {
    fmt.Println(i, v)
  }
}
```

### 5.2.2 Slice - Slicing

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3, 4, 5}
  fmt.Println(x)

  // [2, 3, 4, 5]
  fmt.Println(x[1:])

  // [2, 3]
  fmt.Println(x[1:3])
}
```

### 5.2.3 Slice - Append

`func append(slice []T, elements ...T) []T` where T is a placeholder for any given type

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3}
  fmt.Println(x)

  x = append(x, 4, 5, 6)
  fmt.Println(x)

  y := []int{7, 8, 9}
  x = append(x, y...)
  fmt.Println(x)
}
```

`...` before the type in function signature: unlimited number of this type (`variadic`)

`...` after the variable: destructure

### 5.2.4 Slice - Delete

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3, 4, 5, 6, 7}
  fmt.Println(x)

  // Delete 4
  x = append(x[:3], x[4:]...)
}
```

### 5.2.5 Slice - Make

Create new slice

- Composite literal
- `make`

When u create a slice, it's sitting on top of an array.

Using `make` when we know the length of the underlying array. Avoiding recreate the underlying array
every time we append new value to the slice.

```go
package main

import "fmt"

func main() {
  x := make([]int, 10, 100)
  fmt.Println(x) // [0 0 0 0 0 0 0 0 0 0]
  fmt.Println(len(x)) // length: 10
  fmt.Println(cap(x)) // capacity: 100 - we have 100 slot in the underlying array to use

  x = append(x, 11)
  fmt.Println(len(x)) // length: 11
  fmt.Println(cap(x)) // capacity: 100 - we have 100 slot in the underlying array to use
}
```

### 5.2.6 Slice - Multidimensional

Thinking like a spreadsheet (slice of slice)

```go
package main

import (
  "fmt"
)

func main() {
  jb := []string{"Hieu", "Ha"}
  fmt.Println(jb)

  mb := []string{"Nguyen", "Nguyen"}
  fmt.Println(mb)

  xp := [][]string{jb, mp}
  fmt.Println(xp)
}
```

### 5.3. Map

Map are key - value store. Super efficient look up.

```go
package main

import "fmt"

func main() {
  m := map[string]int{
    "Hieu": 1103,
    "Ha": 2010
  }

  fmt.Println(m)
  fmt.Println(m["Hieu"]) // 1103
  fmt.Println(m["unknown"]) // zero value of int 0

  v, ok = m["unknown"]
  fmt.Println(v)
  fmt.Println(ok)

  if !ok {
    fmt.Println("Key not found")
  }

  // Or
  if v, ok = m["unknow"], ok {
    fmt.Println("Exist: ", v)
  }
}
```

#### 5.3.1. Map - Add Element

```go
m["new"] = 1
```

#### 5.3.2. Map - Range

```go
for k, v in range m {
  fmt.Println(k, v)
}
```

#### 5.3.3. Map - Delete an entry

```go
delete(m, "key")
```

## 6. Structs

A struct is a data structure which allows us to compose together values of different types

```go
package main

import (
  "fmt"
)

type person struct {
  first string
  last string
}

func main() {
  // We create a value of type person
  p1 := person{
    first: "Hieu",
    last: "Nguyen"
  }

  // We create another value of type person
  p2 := person{
    first: "Ha",
    last: "Nguyen"
  }

  fmt.Println(p1)
  fmt.Println(p2)

  fmt.Println(p1.first, p1.last)
  fmt.Println(p2.first, p2.last)
}
```

### 6.1. Embedded Structs

```go
package main

import (
  "fmt"
)

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  kill int
}

func main() {
  sa1 := secretAgent{
    kill: 0,
    person: person{
      first: "Hieu",
      last: "Nguyen"
    }
  }

  fmt.Println(sa1)
  fmt.Println(sa1.first, sa1.last, sa1.kill)

  // When there is name collision
  fmt.Println(sa1.person.first, sa1.first)
}
```

### 6.2. Reading documentation

A struct is a sequence of named elements, call `fields`, each of which has a `name` and a `type`.

### 6.3. Anonymous Struct

```go
package main

import (
  "fmt"
)

func main() {
  p1 := struct {
    first, last string
  }{
    first: "Hieu",
    last: "Nguyen"
  }

  fmt.Println(p1)
}
```

## 7. Functions

All about modular. **Everything in Go is PASS BY VALUE**

I'm going to take my code and break it up into small modular chunks, small chunks, little modules.

Function signature

`func (r receiver) identifier(parameters) (return(s)) { code }`

```go
func foo() {
  fmt.Println("Hello")
}

func bar(s string) {
  fmt.Println("Hello", s)
}

func woo(s string) string {
  return fmt.Sprint("Hello from, ", s)
}

func mouse(firstname, lastname string) (string, bool) {
  a := "Hello"
  b := true

  return a, b
}

x,y := mouse("Hieu", "Nguyen")
```

### 7.1. Variadic parameter

If function `f` is `variadic` with a `final paramter` `p` of type `...T`, then within `f`
the type of `p` is `equivalent to type []T`. If `f` is invoked with `no actual arguments for p`,
the value passed to `p` is `nil`. (Language spec)

func (r receiver) identifier(paramter(s)) (return(s)) {

}

Define variadic paramter.

Define function `sum` which take an unlimitted numbers of paramters type `int`

```go
func sum(x ...int) int {
  fmt.Println(x)
  fmt.Printf("%T\n", x)

  s := 0

  for i, v := range x {
    sum += v
  }

  return sum
}

x := sum(1, 2, 3, 4, 5)
// [1 2 3 4 5]
// []int | slice of int
```

### 7.2. Unfurling a slice (destructuring a slice)

```go
package main

import (
  "fmt"
)

func main() {
  xi := []int{1, 2, 3, 4, 5}

  sum(xi) // Wrong
  sum(xi...) // Correct

  sum() // Correct
}

func sum(numbers ...int) int {
  sum := 0

  for _, v := range numbers {
    sum += v
  }

  return sum
}
```

### 7.3. Defer

This is where we really get to start to take our program and decouple it.
Break it up into pieces and be able to start modularize it.

A "defer" statement invokes a function whose execution is deferred to the moment the `surrounding function`
returns, either because the surrounding function executed a `return statement`, reached the end of its
`function body`, or because the corresponding goroutine is `panicking`.

```go
package main

import "fmt"

func main() {
  foo()
  bar()
}

func foo() {
  fmt.Println("foo")
}

func bar() {
  fmt.Println("bar")
}
```

As normal, `foo` will be invoked first, and then `bar`.
However, if the `defer` `foo`, then `bar` will be executed first and then `foo`

```go
package main

import "fmt"

func main() {
  defer foo()
  bar()
}

func foo() {
  fmt.Println("foo")
}

func bar() {
  fmt.Println("bar")
}
```

Usage:

I open a file in my program. I want to make sure that that file gets closed when I'm done with it,
right after I open it, I could say `defer` and I can run a `function` to `close that file`.

Defer is always going to run whenever the containing function exists.

### 7.4. Methods

Attach methods to struct.

Function signature: `func (r receiver) identifier(paramter(s)) (return(s)) {}`

```go
package main

import (
  "fmt"
)

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  kill int
}

func (s secretAgent) speak() { // method speak attached to type secretAgent, have accessed to all fields of this type
  fmt.Println("I am, ", s.first, s.last)
}

func speak(s secretAgent) {  } // function speak with paramter of type secretAgent

func main() {
  sa1 := secretAgent {
    person: person {
      first: "Hieu",
      last: "Nguyen"
    },
    kill: 0
  }

  sa1.speak()

  sa2 := secretAgent {
    person: { "Ha", "Nguyen" },
    kill: 0
  }

  sa2.speak()
}
```

### 7.5. Interfaces & Polymorphism (search for Bill kenedy composition)

Interfaces allow us to define behaviour and to do polymorphism.

Any body who has `any type` that has this method `speak` is also of type `human`.

A value can be of more than one type. `sa1` can be type of `secretAgent` or `human`.

```go
package main

import (
  "fmt"
)

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  kill int
}

/**
* Any body who has any type that has this method `speak` is also of type human
*/
type human interface {
  speak()
}

func (s secretAgent) speak() { // method speak attached to type secretAgent, have accessed to all fields of this type
  fmt.Println("I am, ", s.first, s.last)
}

func speak(s secretAgent) {  } // function speak with paramter of type secretAgent

func main() {
  sa1 := secretAgent {
    person: person {
      first: "Hieu",
      last: "Nguyen"
    },
    kill: 0
  }

  sa1.speak()

  sa2 := secretAgent {
    person: { "Ha", "Nguyen" },
    kill: 0
  }

  sa2.speak()

  bar(sa2)
}

func bar(h human) {
  h.speak()
}
```

If an `interface` is `empty`, that means every type will implement this `interface` by default.
So `fmt.Println(a ...interface{})` will accept unlimitted arguments of any type.

### 7.6. Assertion

```go
package main

import "fmt"

func main() {

}

func SwitchOnType(x interface{}) {
    switch x.(type) {
      case int:
        fmt.Println("int")
        fallthrough
      case string:
        fmt.Println("string")
        fallthrough
      case contact:
        fmt.Println("contact struct")
    }
}
```

Asserting a type to be more specific type.
Conversion `interface` to the `concrete type`

```go
type human interface {
  speak()
}

func foo(h human) {
  switch h.(type) {
    case person:
      fmt.Println("Fullname: ", h.(person).first, h.(person).last)

    case secretAgent:
      fmt.Println("Kill: ", h.(secretAgent).kill)
  }
}
```

### 7.7. Anonymous function

Anonymous function signature: `func() {}`

`IIFE` func

```go
package main

import "fmt"

func main() {
  func(a int, b int) {
    fmt.Println("Hello From Anoy", a, b)
  }(1, 2)
}
```

### 7.8. Function expression (Assign a func to variable)

Expression is something that comes down to a value.
And the value that we're going to have here is a function,
and we gonna assign it to a variable.

In Go, `function is first-class citizens`. Function can be used as any other variables,
any other types.

```go
package main

import (
  "fmt"
)

func main() {
  f := func() {
    fmt.Println("my first func expression")
  }

  f()

  f = func(x int) {
    fmt.Println(x)
  }

  f(1)
}
```

### 7.9. Return a func

Function `foo` return a `string` value

Function `bar` return a `func` with the following signature: `func() int`

```go
package main

import "fmt"

func main() {
  f := bar();
  r := f();

  fmt.Println(r)
  fmt.Printf("%T\n", f) // func() int
}

func foo() string {
  s := "Hello"
  return s
}

func bar() func() int {
  return func() int {
    return 11;
  }
}
```
