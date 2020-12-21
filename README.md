# lang

Simple language detection library written on pure Go.

## Summary

* Require Go version >= 1.12
* Written on pure Go
* Supported languages: German (de), Greek (el), English (en), Spanish (es), French (fr), Italian (it), Georgian (ka), Russian (ru)
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
}
```
