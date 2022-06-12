package mul

import "testing"

func TestMul(t *testing.T) {
	if o := Mul(1, 1); o != 1 {
		t.Errorf("Mul(1,1) != %d", o)
	}
}
func TestMul1(t *testing.T) {
	if o := Mul1(1, 1); o != 1 {
		t.Errorf("Mul(1,1) != %d", o)
	}
}
