package verify_pack

import (
	"math"
	"testing"

	a "github.com/devskiller/task/calculator"
)

func TestAddition_Assessment(t *testing.T) {
	total := a.Add(-1, -1)
	if total != -2 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, -2)
	}
}

func TestSubstraction_Assessment(t *testing.T) {
	total := a.Subtract(2, -1)
	if total != 3 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 3)
	}
}

func TestMultiplication_Assessment(t *testing.T) {
	total := a.Multiply(2, 0)
	if total != 0 {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 0)
	}
}

func TestDivision_Assessment(t *testing.T) {
	total := a.Divide(4, 0.0)
	if total != math.Inf(1) {
		t.Errorf("add result was incorrect, got: %v, want: %v.", total, 2)
	}
}
