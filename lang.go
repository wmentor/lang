package lang

import (
	"io"

	"github.com/wmentor/lang/internal/lang"
)

const (
	UnknownLang string = "??"
)

func Detect(in io.Reader) string {
	return lang.Detect(in)
}

type Detector struct{}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) Detect(in io.Reader) string {
	return Detect(in)
}
