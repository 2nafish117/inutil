package tests

import (
	"inutil/lexer"
	"testing"
)

var keywordCases = []TrueFalseCase{
	{`_hey`, false},
	{`hmmm`, false},
	{`is_yes`, false},
	{`isOk`, false},
	{`fortyNine547_`, false},
	{`3156`, false},
	{`47bananas_`, false},

	{`var`, true},
	{`func`, true},
	{`struct`, true},
	{`type`, true},
	{`return`, true},
	{`if`, true},
	{`else`, true},
	{`for`, true},
}

var keywords = []string{
	"var",
	"func",
	"struct",
	"type",
	"return",
	"if",
	"else",
	"for",
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TestKeyword(t *testing.T) {

	for i, c := range keywordCases {
		l := lexer.NewLexer([]byte(c.test))

		tok := l.Scan()
		got := contains(keywords, tok.IDValue())

		if got == c.expected {
			t.Logf("TEST PASSED %d: (%s)", i, c.test)
		} else {
			t.Errorf("TEST FAILED %d: (%s): expected %v, got %v", i, c.test, c.expected, got)
		}
	}
}
