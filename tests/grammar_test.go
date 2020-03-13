package tests

import (
	"inutil/lexer"
	"inutil/parser"
	"testing"
)

var grammarCases = []TrueFalseCase{
	{`

	// this is comment
	func lol(hehe, hehem, hhh) {
		a = 2;
		b = a + 2.5 + v;
		bb = 67;
	}

	func id() {}

	func id() {
		{
			d = 6.7 == 56;
			hey = (g <= there() || 6.78) && 6.7;
		}

		f = -g;
		hel = -(s + g + 4 - 5 / 678);
		a = -4 * -5;
	}

	func main() 
	{
		lol(4, 6, 8, 8);
		hey = 67;
		s = there()
		;

		if hello == 4 {
			bring(ye, wassup + 56, 7.8 - 56); 
			it(4.67 || 67);
			on(1, 4, 7);
		}

		if what >= 56 {
			do(); something();
		} else {
			hi = hi + 45;
		}

		for ; 1 ; {

		}

		for i = 0 ; i < 78 ; i = i + 56 
		{
			// stuffhere(bim, 67.4, 2.8, 9);
			// g = 78;
			if d > 0 {
				more_stuff();
			}

			for ;; {
				beep();
				for ;; {

				}
			}
		}
	}`, true},
}

func TestGrammar(t *testing.T) {
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
