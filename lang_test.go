package lang

import (
	"strings"
	"testing"
)

func TestLang(t *testing.T) {

	tD := func(txt string, lang string) {
		if Detect(strings.NewReader(txt)) != lang {
			t.Fatalf("Detect lang failed for: %s", txt)
		}
	}

	tD("123 12341234 5243245 324534", UnknownLang)
	tD("Hello, World!", "en")
	tD("What do you about that?", "en")
	tD("Привет, мир!", "ru")
	tD("Частостный словарь русского языка", "ru")
}
