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