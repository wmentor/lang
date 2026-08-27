//go:build ignore

package main

import (
	"strings"

	"github.com/wmentor/lang"
)

func main() {
	println(lang.Detect(strings.NewReader("Hello, world!")))             // en
	println(lang.Detect(strings.NewReader("Привет, мир!")))              // ru
	println(lang.Detect(strings.NewReader("Hallo Welt!")))               // de
	println(lang.Detect(strings.NewReader("Բարեւ աշխարհ!")))             // hy
	println(lang.Detect(strings.NewReader("你好世界")))                      // zh
	println(lang.Detect(strings.NewReader("مرحبا بالعالم!")))            // ar
	println(lang.Detect(strings.NewReader("გამარჯობა მსოფლიო!")))        // ka
	println(lang.Detect(strings.NewReader("Ciao, mondo meraviglioso!"))) // it
	println(lang.Detect(strings.NewReader("Bonjour le monde!")))         // fr
	println(lang.Detect(strings.NewReader("Saluton Mondo!")))            // es
	println(lang.Detect(strings.NewReader("Γεια σου Κόσμο!")))           // el
}
