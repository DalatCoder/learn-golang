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
