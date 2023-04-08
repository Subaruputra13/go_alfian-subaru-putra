package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}

	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstrac(t *testing.T) {
	if Substrac(10, 5) != 5 {
		t.Error("Expected 10 (-) 5 to equal 5")
	}
}

func TestMult(t *testing.T) {
	if Mult(5, 5) != 25 {
		t.Error("Expected 5 (*) 5 to equal 25")
	}
}

func TestDivi(t *testing.T) {
	if Divi(10, 5) != 2 {
		t.Error("Expected 10 (/) 5 to equal 2")
	}
}
