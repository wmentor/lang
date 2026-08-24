package lang_test

import (
	"strings"
	"testing"

	"github.com/wmentor/lang/internal/lang"
)

func TestLang(t *testing.T) {
	t.Parallel()

	tD := func(txt string, waitLang string) {
		if res := lang.Detect(strings.NewReader(txt)); res != waitLang {
			t.Fatalf("Detect lang failed for: %s expect=%s res=%s", txt, waitLang, res)
		}
	}

	tD("123 12341234 5243245 324534", lang.UnknownLang)
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
	tD("Բարեւ աշխարհ!", "hy")
	tD("مرحبا بالعالم!", "ar")
	tD("السلام عليكم", "ar")
	tD("تعلم اللغة العربية بالكمبيوتر", "ar")
	tD("كل شيء هادئ في بغداد", "ar")
	tD("你好世界", "zh")
	tD("t薩克斯卷入马可尼丑闻（Marconi scandal）。事件中，两人被懷疑收受马可尼公司的利益", "zh")
	tD("1917年", "zh")
	tD("兼顧了快速原型的迭代特徵以及瀑布模型的系統化與嚴格監控", "zh")
	tD("随后《全国观察家》发表了威尔斯关于时间旅行的设想的连载文章", "zh")
}

func TestCompare(t *testing.T) {
	t.Parallel()

	tC := func(t1 string, t2 string, wait float64) {
		val := lang.Compare(strings.NewReader(t1), strings.NewReader(t2))
		vi := int64(val * 1000000)
		wi := int64(wait * 1000000)
		if vi != wi {
			t.Fatalf("Compare v=%d w=%d text1=%s text2=%s", vi, wi, t1, t2)
		}
	}

	tC("привет, мир!", "Привет Мир!", 1)
	tC("привет, мир!", "Мир фэнтези", 0.095238)

	t1 := `Аналитический сервис TextFrame предоставляет HTTP интерфейс для анализа русскоязычного контента.
	В качестве формата сообщений API используется JSON. С помощью API можно определить язык контента, проверить
	наличие нецензурной лексики, определить вектор вероятности категорий контента`

	t2 := `В качестве формата сообщений API используется JSON. С помощью API можно определить язык контента,
	проверить наличие нецензурной лексики, определить вектор вероятности категорий контента. Кроме этого анализатор
	возвращает список наиболее часто встречающихся ключевых слов, организаций, брендов, людей, стран, географических точек.`

	tC(t1, t2, 0.452127)
}
