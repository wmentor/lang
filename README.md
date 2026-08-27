# lang

![test](https://github.com/wmentor/lang/workflows/test/badge.svg)
[![https://pkg.go.dev/github.com/wmentor/lang](https://pkg.go.dev/badge/github.com/wmentor/lang.svg)](https://pkg.go.dev/github.com/wmentor/lang)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple language detection library written on pure Go.

## Summary

* Require Go version >= 1.18
* Written on pure Go
* Supported languages: Arabic(ar), Armenian (hy), Chinese(zh), German (de), English (en), French (fr), Georgian (ka), Greek (el), Italian (it), Russian (ru), Spanish (es).
* No external dependencies
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
	println(lang.Detect(strings.NewReader("Hello, world!")))             // en
	println(lang.Detect(strings.NewReader("Привет, мир!")))              // ru
	println(lang.Detect(strings.NewReader("Hallo Welt!")))               // de
	println(lang.Detect(strings.NewReader("Բարեւ աշխարհ!")))            // hy
	println(lang.Detect(strings.NewReader("你好世界")))                   // zh
	println(lang.Detect(strings.NewReader("مرحبا بالعالم!")))            // ar
	println(lang.Detect(strings.NewReader("გამარჯობა მსოფლიო!")))        // ka
	println(lang.Detect(strings.NewReader("Ciao, mondo meraviglioso!"))) // it
	println(lang.Detect(strings.NewReader("Bonjour le monde!")))         // fr
	println(lang.Detect(strings.NewReader("Saluton Mondo!")))            // es
	println(lang.Detect(strings.NewReader("Γεια σου Κόσμο!")))           // el
}
```
