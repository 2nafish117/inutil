package tests

import (
	"inutil/lexer"
	"inutil/parser"
	"testing"
)

var forCases = []TrueFalseCase{
	{`for ;; {}`, true},
	{`for ;; { do(); stuff(); here(); }`, true},
	{`for a = 1; a < 56; a= a + b {}`, true},
	{`for ; a ; {}`, true},
	{`for ; foon() < 5 ; {}`, true},
	{`for ; anotherVar < 5 ; {}`, true},
	{`for i = 0; i < 6 ;i = i+ 2 {
		for j = 0;j < 56;j = j + 56 
		{		
			for p = 0;j < 56;j = j + 56 
			{
				for j = 0;j < 56;j = j + 56 
				{
					do_stuff();
					b = 9;
					f = 67;
					if d > 56 {
						hello = do_another_thing();
					}
				}
			}
		}
	}`, true},

	{`for ; a {}`, false},
	{`for ; a ; {}`, false},
	{`for b = 0; b < 5; b = b + 1 hy(); `, false},
}

func TestFor(t *testing.T) {
	for _, c := range funcCases {
		l := lexer.NewLexer([]byte(c.test))
		p := parser.NewParser()
		ret, err := p.Parse(l)

		t.Log(ret, err)
		got := (err == nil)

		if got != c.expected {
			t.Errorf("TEST FAILED: (%s): expected %v, got %v", string(c.test), nil, err)
		}
	}
}
