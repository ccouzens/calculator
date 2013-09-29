package calculator

import "testing"

func TestLexLiteral(t *testing.T) {
	const in = "45"
	out1 := Token{in, t_literal_number}
	ch := make(chan Token)
	go Lex(in, ch)
	x := <-ch
	if x != out1 {
		t.Errorf("Lex(%v) = %v, want %v", in, x, out1)
	}
}

func TestLexWhiteSpace(t *testing.T) {
	const in = " \t"
	var out1 = Token{in, t_white_space}
	ch := make(chan Token)
	go Lex(in, ch)
	x := <-ch
	if x != out1 {
		t.Errorf("Lex(%v) = %v, want %v", in, x, out1)
	}
}

func TestLexPlus(t *testing.T) {
	const in = "+"
	var out1 = Token{in, t_plus}
	ch := make(chan Token)
	go Lex(in, ch)
	x := <-ch
	if x != out1 {
		t.Errorf("Lex(%v) = %v, want %v", in, x, out1)
	}
}

func TestLexCompound(t *testing.T) {
	const in = "3 + 28"
	var out = []Token{
		{"3", t_literal_number},
		{" ", t_white_space},
		{"+", t_plus},
		{" ", t_white_space},
		{"28", t_literal_number},
	}

	const length = 5
	ch := make(chan Token)
	go Lex(in, ch)

	for i := 0; i < length; i++ {
		x, ok := <-ch
		if !ok {
			t.Errorf("Expected Lex(%v)[%d] to exist", in, i)
			return
		} else if x != out[i] {
			t.Errorf("Lex(%v)[%d] = %v, want %v", in, i, x, out[i])
			return
		}

	}
}

func TestParseLiteral(t *testing.T) {
	const in, out = "45", 45
	x, err := Parse(in)
	if err != nil {
		t.Errorf("Parse(%v) raised %v", in, err)
	} else if x != out {
		t.Errorf("Parse(%v) = %v, want %v", in, x, out)
	}
}

func TestParseAddition(t *testing.T) {
	const in, out = "3 + 5", 8
	x, err := Parse(in)
	if err != nil {
		t.Errorf("Parse(%v) raised %v", in, err)
	} else if x != out {
		t.Errorf("Parse(%v) = %v, want %v", in, x, out)
	}
}
