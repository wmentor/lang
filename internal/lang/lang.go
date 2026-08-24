package lang

import (
	"bufio"
	"embed"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/wmentor/lang/internal/mcounter"
)

const (
	UnknownLang = "??"

	trgmLimit    = 1000
	compareAlloc = 1024
)

//go:embed *.lng
var fs embed.FS

var (
	lngData map[uint64]string
)

func init() { //nolint:gochecknoinits // ok.
	lngData = make(map[uint64]string, trgmLimit*30)

	files, err := fs.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		lng := file.Name()[:2]
		loadLang(lng, file.Name())
	}
}

func loadLang(lng string, filename string) {
	rf, err := fs.Open(filename)
	if err != nil {
		panic("load file " + filename + " error: " + err.Error())
	}
	defer rf.Close()

	list := make([]uint64, trgmLimit)

	if err = binary.Read(rf, binary.BigEndian, &list); err != nil {
		panic("decode file " + filename + " error: " + err.Error())
	}

	for _, trgm := range list {
		if v, has := lngData[trgm]; has {
			lngData[trgm] = v + " " + lng
		} else {
			lngData[trgm] = lng
		}
	}
}

func TrgmMap(in io.Reader, ret map[uint64]uint64) {
	br := bufio.NewReader(in)

	st0 := uint64('_')
	st1 := uint64(0)
	st2 := uint64(0)
	st3 := uint64(0)

	pushRune := func(r rune) {
		st3 = (st2 << 16) | uint64(r)
		st2 = (st1 << 16) | uint64(r)
		st1 = (st0 << 16) | uint64(r)
		st0 = uint64(r)
		if st3 != 0 {
			ret[st3]++
		}
	}

	lastSpace := true

	for {
		r, _, err := br.ReadRune()
		if err != nil && r == 0 {
			break
		}

		if r = unicode.ToLower(r); r == 'ё' {
			r = 'е'
		}

		if unicode.IsLetter(r) { //nolint:nestif // skip.
			if unicode.Is(unicode.Han, r) {
				ret[uint64(r)]++
				if !lastSpace {
					lastSpace = true
					pushRune('_')
				}
			} else {
				lastSpace = false
				pushRune(r)
			}
		} else {
			if !lastSpace {
				lastSpace = true
				pushRune('_')
			}
		}
	}

	if !lastSpace {
		pushRune('_')
	}
}

func TrgmTop(hash map[uint64]uint64) []uint64 {
	list := make([]uint64, len(hash))
	i := 0
	for k := range hash {
		list[i] = k
		i++
	}

	sort.Slice(list, func(i, j int) bool {
		vi := hash[list[i]]
		vj := hash[list[j]]
		return vi > vj || vi == vj && list[i] < list[j]
	})

	if len(list) > trgmLimit {
		list = list[:trgmLimit]
	}

	return list
}

func TrgmSave(filename string, list []uint64) error {
	wh, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file %q error: %w", filename, err)
	}
	defer wh.Close()

	fmt.Printf("save %d trgms to file %s\n", len(list), filename)

	if err = binary.Write(wh, binary.BigEndian, list); err != nil {
		return fmt.Errorf("binary write to file %q error: %w", filename, err)
	}

	return nil
}

func Detect(in io.Reader) string {
	lns := mcounter.New()

	hash := make(map[uint64]uint64)
	TrgmMap(in, hash)

	for k, v := range hash {
		for _, l := range strings.Fields(lngData[k]) {
			lns.Inc(l, uint64(v)) //nolint:unconvert // ok.
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

func Compare(in1 io.Reader, in2 io.Reader) float64 {
	h1 := make(map[uint64]uint64, compareAlloc)
	h2 := make(map[uint64]uint64, compareAlloc)

	var wg sync.WaitGroup

	process := func(in io.Reader, h map[uint64]uint64) {
		defer wg.Done()
		TrgmMap(in, h)
	}

	wg.Add(2)
	go process(in1, h1)
	go process(in2, h2)

	wg.Wait()

	cnt := 0

	for k := range h2 {
		if v := h1[k]; v > 0 {
			cnt++
		}
		h1[k] = 1
	}

	if len(h1) == 0 {
		return 1
	}

	return float64(cnt) / float64(len(h1))
}
