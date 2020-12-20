// +build ignore

package main

import (
	"github.com/wmentor/embed"
)

func main() {

	for _, l := range []string{"en", "ru"} {
		src := l + ".txt"
		res := src + ".go"
		err := embed.Make(src, res, "data", "github.com/wmentor/lang/data/"+src)
		if err != nil {
			panic(err)
		}
	}

}
