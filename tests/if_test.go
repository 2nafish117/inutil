package tests

import (
	"inutil/lexer"
	"inutil/parser"
	"testing"
)

var ifCases = []TrueFalseCase{
	{`if c {}`, true},
	{`if c {
		if not_true_func() {
			do_stuff(lolis, are, gay);
			bread = get_bread();
			cofee = bread + ree;
			// a coorect comment
			/*
				miltilines yeahh!!
			*/
		}
	}`, true},
	{`if anotherOneBitesThedust {
		for i = 0;i < 455; i= i * i {
			one = another_one(4, 6, 8);
			print(one);
		}
	}`, true},

	{`if {}`, false},
	{`if (c) {}`, false},
	{`if () {}`, false},
	{`if 4 }`, false},
	{`if 4 {
		if in_contact() {
			// justa random comment
		}
	}`, false},
}

func TestIf(t *testing.T) {
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
