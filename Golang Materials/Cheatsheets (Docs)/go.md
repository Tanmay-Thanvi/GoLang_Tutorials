---
title: Go
layout: 2017/sheet
prism_languages: [go, bash]
weight: -3
tags: [Featured]
category: C-like
updated: 2020-06-21
---
## Getting started

{: .-three-column}

### Introduction

{: .-intro}

- [A tour of Go](https://tour.golang.org/welcome/1) _(tour.golang.org)_
- [Go repl](https://repl.it/languages/go) _(repl.it)_
- [Golang wiki](https://github.com/golang/go/wiki/) _(github.com)_

### Hello world

{: .-prime}

#### hello.go

{: .-file}

```go
package main

import "fmt"

func main() {
  message := greetMe("world")
  fmt.Println(message)
}

func greetMe(name string) string {
  return "Hello, " + name + "!"
}
```

```bash
$ go build
```

Or try it out in the [Go repl](https://repl.it/languages/go), or [A Tour of Go](https://tour.golang.org/welcome/1).

### Variables

#### Variable declaration

```go
var msg string
var msg = "Hello, world!"
var msg string = "Hello, world!"
var x, y int
var x, y int = 1, 2
var x, msg = 1, "Hello, world!"
msg = "Hello"
```

#### Declaration list

```go
var (
  x int
  y = 20
  z int = 30
  d, e = 40, "Hello"
  f, g string
)
```

#### Shortcut of above (Infers type)

```go
msg := "Hello"
x, msg := 1, "Hello"
```

### Constants

```go
const Phi = 1.618
const Size int64 = 1024
const x, y = 1, 2
const (
  Pi = 3.14
  E  = 2.718
)
const (
  Sunday = iota
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)
```

Constants can be character, string, boolean, or numeric values.

See: [Constants](https://tour.golang.org/basics/15)

## Basic types

{: .-three-column}

### Strings

```go
str := "Hello"
```

```go
str := `Multiline
string`
```

Strings are of type `string`.

### Numbers

#### Typical types

```go
num := 3          // int
num := 3.         // float64
num := 3 + 4i     // complex128
num := byte('a')  // byte (alias for uint8)
```

#### Other types

```go
var u uint = 7        // uint (unsigned)
var p float32 = 22.7  // 32-bit float
```

### Arrays

```go
// var numbers [5]int
numbers := [...]int{0, 0, 0, 0, 0}
```

Arrays have a fixed size.

### Slices

```go
slice := []int{2, 3, 4}
```

```go
slice := []byte("Hello")
```

Slices have a dynamic size, unlike arrays.

### Pointers

```go
func main () {
  b := *getPointer()
  fmt.Println("Value is", b)
}
```
