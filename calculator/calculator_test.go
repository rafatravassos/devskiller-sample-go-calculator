package calculator

import "testing"

func TestAddition(t *testing.T) {
	total := Add(2, 3)
	if total != 5 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 5)
	}
}

func TestSubstraction(t *testing.T) {
	total := Subtract(2, 3)
	if total != -1 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, -1)
	}
}

func TestMultiplication(t *testing.T) {
	total := Multiply(2, 3)
	if total != 6 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 6)
	}
}

func TestDivision(t *testing.T) {
	total := Divide(4, 2)
	if total != 2 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 2)
	}
}
