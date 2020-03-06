package tests

import (
	"inutil/lexer"
	"inutil/token"
	"testing"
)

type TrueFalseCase struct {
	test     string
	expected bool
}

var identifierCases = []TrueFalseCase{
	{`_hey`, true},
	{`hmmm`, true},
	{`is_yes`, true},
	{`isOk`, true},
	{`fortyNine547_`, true},

	{`3156`, false},
	{`47bananas_`, false},
	{`var`, false},
	{`func`, false},
	{`struct`, false},
	{`if`, false},
	{`for`, false},
}

func TestIdentifier(t *testing.T) {

	for i, c := range identifierCases {
		l := lexer.NewLexer([]byte(c.test))

		got := (l.Scan().Type == token.TokMap.Type("identifier"))

		if got == c.expected {
			t.Logf("TEST PASSED %d: (%s)", i, c.test)
		} else {
			t.Errorf("TEST FAILED %d: (%s): expected %v, got %v", i, c.test, c.expected, got)
		}
	}
}
