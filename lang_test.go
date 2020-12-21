package lang

import (
	"strings"
	"testing"
)

func TestLang(t *testing.T) {

	tD := func(txt string, lang string) {
		if res := Detect(strings.NewReader(txt)); res != lang {
			t.Fatalf("Detect lang failed for: %s expect=%s res=%s", txt, lang, res)
		}
	}

	tD("123 12341234 5243245 324534", UnknownLang)
	tD("Brave new World!", "en")
	tD("What do you think about that?", "en")
	tD("Привет, мир!", "ru")
	tD("Частостный словарь русского языка", "ru")
	tD("В Багдаде все спокойно", "ru")
	tD("Hallo Welt!", "de")
	tD("Ciao, mio ​​migliore amico!", "it")
	tD("¡Hola mi mejor amiga en este maravilloso día!", "es")
	tD("Bonjour mon meilleur ami en cette merveilleuse journée!", "fr")
	tD("Bonjour mon jeune ami!", "fr")
	tD("Never gonna give you up", "en")
	tD("გამარჯობა ჩემო ახალგაზრდა მეგობარო!", "ka")
	tD("Γεια σας, ο μικρός μου φίλος!", "el")
}
