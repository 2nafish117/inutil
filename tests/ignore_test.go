package tests

import (
	"inutil/lexer"
	"inutil/token"
	"testing"
)

var ignoreCases = []TrueFalseCase{
	{`// this is a line comment`, true},
	{`/* this is a multi line comment in one line */`, true},
	{`// this is /another line/ comment\ hello world`, true},
	{`/* this is multiline 
		in multiple 
		lines */`, true},
	{` `, true},
	{`	`, true},
	{`
	`, true},
	{` 	
	`, true},
}

func TestIgnore(t *testing.T) {
	for i, c := range ignoreCases {
		l := lexer.NewLexer([]byte(c.test))

		tok := l.Scan()
		got := tok.Type == token.TokMap.Type("INVALID") || tok.Type == token.TokMap.Type("$")

		if got == c.expected {
			t.Logf("TEST PASSED %d: (%s)", i, c.test)
		} else {
			t.Errorf("TEST FAILED %d: (%s): expected %v, got %v", i, c.test, c.expected, got)
		}
	}
}
