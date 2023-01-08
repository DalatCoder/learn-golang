# Go Bootcamp: Master Golang with 1000+ Exercises and Projects

[Udemy](https://udemy.com/course/learn-go-the-complete-bootcamp-course-golang/)

## Basics

### Introduction to Variables

> Declaration syntax: You need to declare a variable before
> you can use it

```go
var n int
```

- `var`: keyword
- `n`: name of the variable, also known as an `identifier`
- `int`: variable's type determines what `type of values` you
  can store in a variable (`strongly-typed`)

There are only two hard things in computer science:

- cache invalidation
- naming things

Names should start with a `letter` or an `underscore`. `Unicode letters` are also OK

```go
package main
import "fmt"

func main() {
    var speed int
    fmt.Println(speed) // 0
}
```

![Image](assets/variables.png)

### Example: `path seperator`

Let's learn how to use multiple result returning expressions
in multiple assignments

`path package` provides utility functions for working with `url path` strings

`Split` function signature

- Input: `path` string
- Output: `directory` and `file`, both have type of `string`

```go
func Split(path string) (dir, file string) {}
```

![Image](assets/pathSeperator.png)

```go
import ( "fmt"; "path" )

func main() {
    var dir, file string
    dir, file = path.Split("css/main.css")

    // discard
    // _, file = path.Split("css/main.css")

    fmt.Println("dir: ", dir)
    fmt.Println("file: ", file)
}
```

Or using `short declaration`: `:=`

```go
import ( "fmt"; "path" )

func main() {
    dir, file := path.Split("css/main.css")

    // discard
    // _, file := path.Split("css/main.css")

    fmt.Println("dir: ", dir)
    fmt.Println("file: ", file)
}
```

### When to use a `short declaration`?

`Short` vs `Normal`

Use `normal declaration`:

- if you don't know the `initial value`
- when you need a `package scoped` variable
- when you want to group variables together for greater readability

```go
package main

// version := 0
var version int

func main() {
    // score := 0 // DONT!
    var score int // already score = 0

    var (
        // related:
        video string

        // closely related
        duration int
        current it
    )
}
```

Use `short declaration`:

- we mostly perfer using `short declaration`
- if you know the `initial value`
- to keep the code concise
- for `redeclaration`
- using inside `if` and `switch` statements to create variables that belong to those statements only

```go
package main

func main() {
    // var width, height = 100, 50 // DONT!
    width, height := 100, 50

    // DONT!
    width = 50 // assigns 50 to width
    color := "red" // new variable: color

    // using redeclaration
    // change width to 50, declare color = red
    width, color := 50, "red"
}
```

### Let's convert a value

GO just wants you to be explicit on everything. That's why it
doesn't let you assign a `float64` to an `int` value.

`Type conversion`: changes the `type` of a value to `another type`

Which value you convert first matters.

Signature: `type(value)`

```go
package main

func main() {
    speed := 100 // int
    force := 2.5 // float64

    speed = int( float64(speed) * force )
}
```

### Example: Get input from terminal

`os package` allows you to access to `operating system functionalities`

`Args` variable belongs to the `os package`

`var Args []string`:

- a slice can store multiple values
- Args's type is a `slice of string`

When you run a Go program, Go puts `the command-line arguments` into `Args` variable automatically

#### Learn the basics of `os.Args

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("%v\n", os.Args)

    fmt.Println("Path:", os.Args[0])
    fmt.Println("1st argument", os.Args[1])
    fmt.Println("2nd argument", os.Args[2])

    fmt.Println("Number of items inside os.Args:", len(os.Args))
}
```

#### Naming things: Recommendations

> There only two hard things in computer science:
> cache invalidation and naming things

- `non-idiomatic` means `not preferred` usage
- `idiomatic` means `preffered` usage

Some common abbreviations used in `Go`

- `var s string`: string
- `var i int`: index
- `var num int`: number
- `var msg string`: message
- `var v string`: value
- `var val string`: value
- `var fv string` flag value
- `var err error`: error value
- `var args []string`: arguments
- `var seen bool`: has seen?
- `var parsed bool`: parsing ok?
- `var buf []byte`: buffer
- `var off int`: offset
- `var op int`: operation
- `var opRead int`: read operation
- `var l int`: length
- `var n int`: number of number of
- `var m int`: another number
- `var c int`: capacity
- `var c int`: character
- `var a int`: array
- `var r rune`: rune
- `var sep string`: separator
- `var src int`: source
- `var dst int`: destination
- `var b byte`: byte
- `var b []byte`: buffer
- `var buf []byte`: buffer
- `var w io.Writeer`: writer
- `var r io.Reader`: reader
- `var pos int`: position

Recommendations

- Use the first few letters of the words: `var fv string // flag value`
- Use fewer letters in smaller scopes (`block` | `function`)

```go
var bytesRead int // number of bytes read // DONT
var n int // number of bytes read
```

- Use the complete words in larger scopes (`package scope`)
- Use `mixedCaps` like this
- Use all `captitals` for acronyms

```go
var localAPI string // DONT
var localApi string // DO
```

- Do not stutter: do not use the same words again and again

```go
var p := player.PlayerScore; // DONT
p := player.Score; // DO
```

- Do not use `under_scores` or `LIKE_THIS`

```golang
const MAX_TIME int // DONT
const MaxTime int // DO
```
