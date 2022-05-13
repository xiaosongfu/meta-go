# Meta-Go

Get app's build go compiler info、build os info、build vcs info and startup time.

## Install

```go
$ go get -d github.com/xiaosongfu/meta-go
```

## Usage

import `github.com/xiaosongfu/meta-go` package in the go file of `main` func located. for example:

```go
package main

import (
	"github.com/xiaosongfu/meta-go"
)

func main() {
	meta.Println()
}
```

if no funcs of this package is used, please use `import  _ "github.com/xiaosongfu/meta-go"` instead.

## APIs

* `func Meta() *META` : Return the `META` instance's pointer.
* `func Text() string` : Return a formatted string of the `META` instance.
* `func Println()` : Print the result of `Text()` func to stdout.
