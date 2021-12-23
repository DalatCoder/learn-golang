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
