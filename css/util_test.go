package css // import "github.com/plazma0/parse/css"

import (
	"testing"

	"github.com/tdewolff/test"
)

func TestIsIdent(t *testing.T) {
	test.That(t, IsIdent([]byte("color")))
	test.That(t, !IsIdent([]byte("4.5")))
}

func TestIsURLUnquoted(t *testing.T) {
	test.That(t, IsURLUnquoted([]byte("http://x")))
	test.That(t, !IsURLUnquoted([]byte(")")))
}

func TestHsl2Rgb(t *testing.T) {
	r, g, b := HSL2RGB(0.0, 1.0, 0.5)
	test.Float(t, r, 1.0)
	test.Float(t, g, 0.0)
	test.Float(t, b, 0.0)

	r, g, b = HSL2RGB(1.0, 1.0, 0.5)
	test.Float(t, r, 1.0)
	test.Float(t, g, 0.0)
	test.Float(t, b, 0.0)

	r, g, b = HSL2RGB(0.66, 0.0, 1.0)
	test.Float(t, r, 1.0)
	test.Float(t, g, 1.0)
	test.Float(t, b, 1.0)
}
