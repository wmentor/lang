# lang

![test](https://github.com/wmentor/lang/workflows/test/badge.svg)
[![https://goreportcard.com/report/github.com/wmentor/lang](https://goreportcard.com/badge/github.com/wmentor/lang)](https://goreportcard.com/report/github.com/wmentor/lang)
[![https://pkg.go.dev/github.com/wmentor/lang](https://pkg.go.dev/badge/github.com/wmentor/lang.svg)](https://pkg.go.dev/github.com/wmentor/lang)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple language detection library written on pure Go.

## Summary

* Require Go version >= 1.18
* Written on pure Go
* Supported languages: Armenian (am), German (de), Greek (el), English (en), Spanish (es), French (fr), Italian (it), Georgian (ka), Russian (ru)
* MIT license

## Install

```plaintext
go get github.com/wmentor/lang
```

## Usage

```golang
package main

import (
  "strings"

  "github.com/wmentor/lang"
)

func main() {
  println(lang.Detect(strings.NewReader("123 1231232332 12"))) // ??
  println(lang.Detect(strings.NewReader("Hello, world!")))     // en
  println(lang.Detect(strings.NewReader("Привет, мир!")))      // ru
  println(lang.Detect(strings.NewReader("Hallo Welt!")))       // de
  println(lang.Detect(strings.NewReader("Բարեւ աշխարհ!")))     // am
}
```
