package tests

import (
	"inutil/lexer"
	"inutil/token"
	"testing"
)

var literalCases = map[string]string{
	`3156`:  "integerLit",
	`-47`:   "integerLit",
	`+6785`: "integerLit",
	`0`:     "integerLit",
	`-0`:    "integerLit",
	`+0`:    "integerLit",

	`46.46`:          "floatLit",
	`-47.58`:         "floatLit",
	`+9725.26897700`: "floatLit",
	`0.0`:            "floatLit",
	`-0.0`:           "floatLit",
	`+0.0`:           "floatLit",
	`.436`:           "floatLit",

	`0..0`: "INVALID",
}

func TestLiteral(t *testing.T) {
	for key, val := range literalCases {
		l := lexer.NewLexer([]byte(key))
		tok := l.Scan()
		got := tok.Type
		expected := token.TokMap.Type(val)
		t.Log(key, val)

		if got != expected {
			t.Errorf("TEST FAILED: (%s): expected %v, got %v", tok.IDValue(), expected, got)
		}
	}
}
