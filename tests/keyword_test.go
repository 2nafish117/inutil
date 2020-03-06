package tests

import (
	"dundi/lexer"
	"dundi/token"
	"testing"
)

var keyword_cases = []Case{
	{`_hey`, "no"},
	{`hmmm`, "no"},
	{`is_yes`, "no"},
	{`isOk`, "no"},
	{`fortyNine547_`, "no"},
	{`3156`, "no"},
	{`47bananas_`, "no"},

	{`var`, "yes"},
	{`func`, "yes"},
	{`struct`, "yes"},
	{`type`, "yes"},
	{`return`, "yes"},
	{`if`, "yes"},
	{`else`, "yes"},
	{`for`, "yes"},
}

func TestKeyword(t *testing.T) {

	for i, c := range keyword_cases {
		l := lexer.NewLexer([]byte(c.test))

		var got string
		if l.Scan().Type == token.TokMap.Type("keyword") {
			got = "yes"
		} else {
			got = "no"
		}

		if got == c.expected {
			t.Logf("TEST %d:\t(%s):\t\t\tPASSED", i, c.test)
		} else {
			t.Errorf("TEST %d:\t(%s):\t\t\tFAILED: expected %s, got %s", i, c.test, c.expected, got)
		}
	}

}
