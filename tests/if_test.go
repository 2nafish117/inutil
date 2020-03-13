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
	{`if something {} else {}`, true},
	{`if something {} else {
		if another {

		} else {
			sayHello(1, 3, 45);
		}
	}`, true},
	{`if something {} else {} 
	if hi {more_stuff();}
	`, true},
	{`if case1 {

	} else {
		if case2 {

		}
	} else {
		if case3 {

		}
	} else {
		//default case
	}`, true},

	{`if something {} else `, false},
	{`if something {} else {} else {}`, false},
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
	for i, c := range funcCases {
		l := lexer.NewLexer([]byte(c.test))
		p := parser.NewParser()
		ret, err := p.Parse(l)

		t.Log(ret, err)
		got := (err == nil)

		if got != c.expected {
			t.Errorf("TEST %d FAILED: (%s): expected %v, got %v", i, string(c.test), nil, err)
		}
	}
}
