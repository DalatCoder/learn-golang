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
