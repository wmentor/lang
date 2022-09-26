package lang

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/wmentor/mcounter"
	ngram "github.com/wmentor/qgram"
)

const (
	UnknownLang string = "??"
)

var (
	Langs []string
	data  map[string]string
)

func init() {
	Langs = []string{"am", "de", "el", "en", "es", "fr", "it", "ka", "ru"}
	data = map[string]string{}

	for _, name := range Langs {
		loadLang(name)
	}
}

func loadLang(name string) {
	in, err := fs.Open(name + ".txt")
	if err != nil {
		panic("unknown data file: " + name)
	}
	defer in.Close()

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
	lns := mcounter.New()

	hash := ngram.CalcMap(in)

	for k, v := range hash {
		for _, l := range strings.Fields(data[k]) {
			lns.Inc(l, uint64(v))
		}
	}

	if len(lns) == 0 {
		return UnknownLang
	}

	list := lns.Slice(1, true)
	if len(list) == 0 {
		return UnknownLang
	}

	return list[0]
}

func Conflicts() {
	for w, ls := range data {
		if list := strings.Fields(ls); len(list) > 1 {
			fmt.Printf("%s %s\n", w, ls)
		}
	}
}
