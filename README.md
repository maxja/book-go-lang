Learning Go
===

Follow up and code examples/exercises from
Book by Jon Bodner

Abstract
---

This is a collection of a follow up notes and code samples taken from the book 
["Learning Go"](https://www.goodreads.com/book/show/55841848-learning-go "on goodreads.com") 
by Jon Bonder.

Версию на русском можно найти на [livelib.ru](https://www.livelib.ru/book/1007447983-go-idiomy-i-patterny-proektirovaniya-dzhon-bodner "на livelib.ru").

All code collected in a tree-formed directory structure where `ch0N` is a short 
for the corresponding chapter.

Every book example code file will have preceding `exmpl_` prefix.

Chapters
---

### Chapter 1. Setting Up Your Go Environment

This chapter all about go toolset and how to setup it.

#### Install the Go Tools

Downloads page at [official Go site](https://go.dev/dl) provide you with 
installers for different OS and CPU architecture.

Author suggest to use [Homebrew](https://brew.sh/) for MacOS and [Chocolatey](https://chocolatey.org/) 
for Windows, as for the Linux most convenient way to use `tag.gz` archive and 
unpack it.

> For the Windows users I suggest to use another package manager, called [Scoop](https://scoop.sh/).

In every OS you need to add go path to user-level or system-level environment 
called PATH.

```
$ go version
```

Will tell the currently used version.

```
$ go env
```

Will list you with all Go related environments. The most crucial at this stage 
are `GOROOT`, place where go lives, and `GOPATH`, where go will store and look 
for dependencies and installed with `go install` binaries.

#### The Go Workspace

Use 
```
$ go install ...
```
to install additional or third-party tools.

By default `GOPATH` points to `$HOME/go` where all source codes lives in a `src` 
directory and all binary in a `bin` folder of it.

Explicitly defined `$GOPATH` is a good idea. It can be done via 
```
export GOPATH=$...
export PATH=$PATH:$GOPATH/bin
```
in Unix-like systems. And via
```
setx GOPATH %...
setx path "%path%;%GOPATH\bin"
```

#### The go Command

Go ships with compiler, formatter, linter, dependency manager, test runner, etc.

##### Run and Build

Code can be run or can be build. Run will compile, run, and wipe build 
artefact on termination.

Author suggest to write "hello world" program and run it.
```
$ go run ./hello.go
```

> Appropriate code sample can be found in `./ch01/hello.go` file.

Build will compile code into binary, by default it will be named same as entry 
point code file or declared package, and can me modified via `-o` flag with 
custom name given.
```
$ go build ./exmpl_hello.go
$ go build -o hello_world ./exmpl_hello.go
```

##### Getting 3rd-Party Go Tools

In Go world there are no centrally hosted package registry, but rather it uses 
source code repository as a reference point for the package.

Any package or go written tool can be installed via
```
$ go install somevcs.registry/author/package@version
```

By setting version to `latest` last available version in this repository will 
be installed.

Dependency manager also will download all package dependency and build project, 
placing binaries to `$GOPATH/bin` directory.

> If path to the package differs from one above, `$GOPROXY` environment 
variable might be used to account those changes.

Author suggest to install `hey` utility to try `go install`.

##### Formatting Your Code

In go world there is no formatting variation, and language authors doesn't left 
any space for dispute about it.

```
$ go fmt ./...
$ gofmt -l -w ./...
```
Will check format of given code, fix what it can, and warn about lines that off.

Author suggest to try `goimports` that extends `go fmt` and sort imports in 
alphabetical order.

```
$ goimports -l -w .
```

> Go fmt won't fix misplaced braces. And that's why: despite that lack of 
semicolon at the end of each command, go lexer will virtually inserts them 
after certain tokens.

##### Linting and Vetting

Author suggest to read [Effective Go](https://go.dev/doc/effective_go) article 
as well as [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) 
page.

Tools like
```
$ golint ./...
```
and
```
$ go vet ./...
```
is used to enforce code style and mitigate inefficiency.

Though, `golangci-lint` suggested to be used as a right tool of collaborative 
work, that insure many aspects of a code regarding code-style and code-quality.

```
$ golangci-lint run ./...
```

Additionally this tool can be tuned via `.golangci.yml` config file placed on 
project level.

#### Choose Your Tools

Any text editor basically, and [Go Playground](https://go.dev/play/).

Go Playground limited, rather then original runtime, but yet useful, to sketch 
something.

In Go Playground you can imitate multi-file structure, by adding 
```
-- filename.go --
```
prior file contents.

#### Makefiles

Author suggest to look at make as an automation tool. But it required to be 
installed upfront.

Example of Makefile can be found in `./ch01/Makefile`.

#### Staying Up to Date

Be cautious while switching between different version of go tools.

### Chapter 2. Primitive Types and Declarations

#### Build-in Types

> `./ch02/exmpl_buildin.go` contains build in types. And by running this 
program, variables will be declared, initialized with type-related zero values 
and are printed.

A [list of go types](https://pkg.go.dev/builtin#pkg-types) could be found in 
original documentation.

#### The Zero Value

Go compiler will initialize uninitialized variables with zero values regarding 
declared types. 

So that variable of `string` type will be empty string, or `int` variable will 
be set to 0.

#### Literals

Literals ≠ variables.

##### Integer number

Integers can be set as a base ten, or other via prefixes: 
* binary `0b`, 
* octal `0o` (`0`),
* hexadecimal `0x`.

Author do not recommend to use `0` prefix, as it might confused other developers.

Underscore symbol can be used for better representation.

##### Floating point

Literals with floating point can be defined as a set on numbers with stop 
delimiter, in exponential form or even in hexadecimal.

Underscore symbol can be used for better representation.

##### Rune

This type of literals represent character and surrounded within single quotes 
`'`. Like so: `'a'`.

Run can be written as an unicode character, 8-bit octal, 8-bit hexadecimal, 
16-bit hexadecimal, or 32-bit unicode number. Also with rune can be defined one 
of a special character.

##### String

String literal can be given withing double quotes `"` around it, or within 
backtick, as a so called raw string. The difference is that string in double 
quotes some of a symbol sequence will be interpreted as non-printable 
characters. 

