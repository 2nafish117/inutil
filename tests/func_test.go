package tests

import (
	"inutil/lexer"
	"inutil/parser"
	"testing"
)

var funcCases = []TrueFalseCase{
	{`func yeet() {}`, true},
	{`func heyy(f, _for, the, fallen) { do(); stuff(); here(); }`, true},
	{`func beat_your_meat(arg1) {}`, true},
	{`func the_cake_is_a_lie(arg1, arg2) {}`, true},
	{`func HighIqFunc(arg1) {}`, true},
	{`func FuncWithStuffIn() {
		a = 5;
		b = 6;
		a = a + b;
		d = a + b;
	}`, true},
	{`func AnotherFuncBitesTheDust() {}`, true},

	{`func 45ThisIsBad(arg1) {}`, false},
	{`nah_i_wont_compile() {}`, false},
	{`func fun()`, false},
}

func TestFunc(t *testing.T) {
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
