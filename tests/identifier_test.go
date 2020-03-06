package tests

import (
	"dundi/lexer"
	"dundi/token"
	"testing"
)

type Case struct {
	test     string
	expected string
}

var identifier_cases = []Case{
	{`_hey`, "yes"},
	{`hmmm`, "yes"},
	{`is_yes`, "yes"},
	{`isOk`, "yes"},
	{`fortyNine547_`, "yes"},

	{`3156`, "no"},
	{`47bananas_`, "no"},
	{`var`, "no"},
	{`func`, "no"},
	{`struct`, "no"},
	{`if`, "no"},
	{`for`, "no"},
}

func TestIdentifier(t *testing.T) {

	for i, c := range identifier_cases {
		l := lexer.NewLexer([]byte(c.test))

		var got string
		if l.Scan().Type == token.TokMap.Type("identifier") {
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
