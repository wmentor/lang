package lang

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/wmentor/embed"
	_ "github.com/wmentor/lang/data"
	"github.com/wmentor/tokens"
)

const (
	UnknownLang string = "??"
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

func Detect(in io.Reader) string {

	lns := map[string]int{}

	tokens.Process(in, func(t string) {
		if v, has := data[t]; has {
			for _, l := range strings.Fields(v) {
				lns[l]++
			}
		}
	})

	if len(lns) == 0 {
		return UnknownLang
	}

	max := 0

	for _, cnt := range lns {
		if max < cnt {
			max = cnt
		}
	}

	var res string
	var threshold int = max / 2

	for l, cnt := range lns {
		if cnt >= threshold {
			if res != "" {
				return UnknownLang
			}
			res = l
		}
	}

	return res
}
