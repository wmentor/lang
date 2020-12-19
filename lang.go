package lang

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/wmentor/embed"
	_ "github.com/wmentor/lang/data"
)

var (
	Langs []string
	data  map[string]string
)

func init() {

	Langs = []string{"en"}
	data = map[string]string{}

	for _, name := range Langs {
		loadLang(name)
	}
}

func loadLang(name string) {
	in, err := embed.Get(fmt.Sprintf("github.com/wmentor/lang/data/%s.txt", name))
	if err != nil {
		return
	}

	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if str = strings.TrimSpace(str); str != "" {
			if v, has := data[str]; has {
				list := strings.Fields(v)
				has = false
				for _, cur := range list {
					if cur == name {
						has = true
						break
					}
				}
				if !has {
					data[str] = v + " " + name
				}
			} else {
				data[str] = name
			}
		}
	}
}
