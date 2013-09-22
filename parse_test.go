package calculator

import "testing"

func TestLiteral(t *testing.T) {
	const in, out = "45", 45
	x, err := Parse(in)
	if err != nil {
		t.Errorf("Parse(%v) raised %v", in, err)
	} else if x != out {
		t.Errorf("Parse(%v) = %v, want %v", in, x, out)
	}
}

func TestAddition(t *testing.T) {
	const in, out = "3 + 5", 8
	x, err := Parse(in)
	if err != nil {
		t.Errorf("Parse(%v) raised %v", in, err)
	} else if x != out {
		t.Errorf("Parse(%v) = %v, want %v", in, x, out)
	}
}
