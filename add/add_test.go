package add

import "testing"

func TestAdd(t *testing.T) {
	if o := Add(1, 1); o != 2 {
		t.Errorf("Add(1,1) != %d", o)
	}
}
func TestAdd1(t *testing.T) {
	if o := Add1(1, 1); o != 2 {
		t.Errorf("Add(1,1) != %d", o)
	}
}
