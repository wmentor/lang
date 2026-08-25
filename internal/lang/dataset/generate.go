//go:build ignore

package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/wmentor/lang/internal/lang"
)

type Key struct {
	Lang string
	Trgm uint64
}

var (
	allTrgms = map[Key]uint64{}
	allLangs = map[string]struct{}{}
	zhTrgms  = map[rune]uint64{}
)

func main() {
	var sourceDir string
	var destDir string
	flag.StringVar(&sourceDir, "source", "", "source data directory")
	flag.StringVar(&destDir, "dest", "", "destination data directory")
	flag.Parse()

	fl, err := os.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	for _, f := range fl {
		if f.IsDir() {
			continue
		}
		if filename := f.Name(); strings.HasSuffix(filename, ".txt.gz") || strings.HasSuffix(filename, ".txt") {
			processFile(sourceDir, filename)
		}
	}

	fmt.Printf("found %d trgms\n", len(allTrgms))

	for curLang := range allLangs {
		saveLang(destDir, curLang)
	}

	fmt.Printf("zh runes: %d\n", len(zhTrgms))
	saveZh(destDir, "zh")
}

func saveLang(dir string, curLang string) {
	filename := filepath.Join(dir, curLang+".lng")

	hash := make(map[uint64]uint64)

	for k, v := range allTrgms {
		if k.Lang == curLang {
			hash[k.Trgm] = v
		}
	}

	if err := lang.TrgmSave(filename, lang.TrgmTop(hash)); err != nil {
		panic("write " + filename + " error: " + err.Error())
	}
}

func saveZh(dir string, curLang string) {
	filename := filepath.Join(dir, curLang+".lng")

	list := make([]uint64, 0, len(zhTrgms))

	for k := range zhTrgms {
		list = append(list, uint64(k))
	}

	sort.Slice(list, func(i, j int) bool {
		ri := rune(list[i])
		rj := rune(list[j])
		vi := zhTrgms[ri]
		vj := zhTrgms[rj]
		return vi > vj || vi == vj && ri < rj
	})

	if err := lang.TrgmSave(filename, list); err != nil {
		panic("write " + filename + " error: " + err.Error())
	}
}

func processFile(dir, filename string) {
	if idx := strings.Index(filename, "_"); idx == 2 {
		fmt.Printf("process %s\n", filename)
		if l := filename[:2]; l != "zh" {
			hash := make(map[uint64]uint64)
			if rh, err := os.Open(filepath.Join(dir, filename)); err == nil {
				defer rh.Close()
				if gz, err := gzip.NewReader(rh); err == nil {
					defer gz.Close()
					allLangs[l] = struct{}{}
					lang.TrgmMap(gz, hash)
					for k, v := range hash {
						allTrgms[Key{l, k}] += v
					}
				}
			}
		} else {
			processZh(filepath.Join(dir, filename))
		}
	}
}

func processZh(filename string) {
	rh, err := os.Open(filename)
	if err != nil {
		return
	}
	defer rh.Close()

	br := bufio.NewReader(rh)

	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(br)
		if err != nil {
			return
		}
		defer gz.Close()
		br = bufio.NewReader(gz)
	}

	for {
		r, _, err := br.ReadRune()
		if err != nil && r == 0 {
			break
		}

		if unicode.IsLetter(r) && unicode.Is(unicode.Han, r) {
			zhTrgms[r]++
		}
	}
}
