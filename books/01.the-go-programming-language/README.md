# The Go Programming Language 

> Go is an open source programming language that makes it easy to build simple, reliable and
> efficient software

Go is especially well suited for building infrastructure

- Networked servers
- Tools
- Systems for programmers

But it is truly a general-purpose language and finds use in domains
as diverse as graphics, mobile applicationis, and maching learning.

## 1. The Origins of Go

The figure below shows the most important influences of ealier programming languages
on the design of Go

![Image](assets/gostructure.png)

go is sometimes described as a `C-like language`, or as `C for the 21st century`

All but the most trivial code examples in the book are avaiable for download form the public Git
repository at `gopl.io`.

Each example is identified by its package import path and may be conventiently featched, built,
and installed using the `go get` command. You'll need to choose a directory to be your `Go workspace`
and set the `GOPATH` environment variable to point to it.

The `go` tool will create the directory if neccessary. 

```sh
export GOPATH=$HOME/gobook
go get gopl.io/ch1/helloworld
$GOPATH/bin/helloworld
```

To run the examples, you will need at least versioin `1.5` of `Go`

```sh
go version
go version go1.5 linux/amd64
```

Follow the instructions at `https://golang.org/doc/install` if the `go` tool on your computer
is older or missing.

## 2. The Origins of Go

The best source for more information about Go is the official web site `https://golang.org`

- Documentation
- `Go Programming Language Specification`
- Standard packages

The `Go blog` at `blog.golang.org` publishes some of the best writing on Go

Go Playground at `play.golang.org`. The playground makes it convenient to perform simple experiments
to check one's understanding of syntax, semantics, or library packages with short programs, and in many ways
takes the place of a `read-eval-print-loop` (REPL) in other languages.

Built atop the Playground, the `Go Tour` at `tour.golang.org` is a sequence of short interactive lessons
on the basic ideas and constructions of `Go`, an orderly walk through the language.

Since `Go` is an open-source project, you can read the code for any type or function in the standard
library online at `https://golang.org/pkg`; the same code is part of the downloaded distribution.

Use this to figure out how something works, or to answer questions about details,
or merely to see how experts write really good `Go`.

## 3. Tutorial

This chapter is a tour of the basic components of `Go`.

### 3.1. Hello, World

```go
package main 

import "fmt"

func main() {
	fmt.PrintLn("Hello")
}
```

`Go` is a compiled language. The `Go` toolchain converts a source program and the things it depends on 
into intrstructions in the native machine language of a computer.

These tools are accessed through a single command called `go` that has a number of subcommands.

- `go run`: 
	
	- compiles the source code from one or more source files whose names end in `.go`
	- links it with libraries
	- runs the resulting `executable` file.

	`go run helloworld.go`

`Go` natively handles `Unicode`.

- `go build`:

	- `go build helloworld.go`
	
	- compiles source
	- save the compiles for later use
	- creates an `executable` binary file
	- can be run with: `./helloworld`

`Go` code is organized into `packages`, which are similar to `libraries` or `modules` in other languages.

A `package` consists of one or more `.go` source files in a `single directory` that define what the `packages` 
does.

Each `source file` begins with a `package declaration`, here `package main`, that states which `package` the
file `belongs to`, followed by `a list` of other `packages` that it `imports`, and then the 
`declarations` of the program that are stored in that file.

The `Go standard library` has over `100 packages` for common tasks like input and output, sorting, and text
manipulation.

`fmp package`

- printing formatted output
- scanning input
- `PrintLn` is one of the basic `output functions` in `fmt`
	
	- Prints one or more values
	- Seperated by spaces
	- Newline character at the end

- `Package main is special`

	- `define` a standalone `executable` program, `not a library`
	- `function main`

		- where the execution of the program begins
		- whatever `main` does is what the program does
		- `main` will normally call upon functions in other packages to do much of the work

`import` declaration

	- what packages are needed by this source file
	
You must import exactly the `packages` you need. A programm `will not` compile if there are missing
imports or if there are unneccessary ones.

This strict requirements prevents references to unused packages from accummulating as programs evolve.

Order 

- `package` declaration
- `package` declaration
- declaration of `functions - func`, `variables - var`, `constants - const`, and `types - type`

A function declaration consists of 
	
- keyword `func`
- the name of the function 
- a parameter list (`empty for main`)
- a result list (`empty for main`)
- the body of the function

`Go` `does not` require `semicolons` at the ends of statements or declarations, `except` where two 
or more `appear on the same line`.

In effect, newlines following certain tokens are converted into `semicolons`, so where newlines
are placed matters to proper parsing of `Go` code.

`Go` takes a strong stance on `code formatting`.

The `gofmt` tool rewrites code into the `standard format`, and the `go` tool's `fmt` subcommand applies
`gofmt` to all the files in the specified `package`, or the ones in the current directory by default.

Many text editors can be configured to run `gofmt` each time you save a file, so that your source code
is always properly formatted. 

A related tool, `goimports`, additionally manages the `insertion` and `removal` of import 
declarations as needed. it is not part of the `standard` distribution but you can obtain it with this command

```sh
go get golang.org/x/tools/cmd/goimports
```