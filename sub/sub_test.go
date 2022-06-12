package sub

import "testing"

func TestSub(t *testing.T) {
	if o := Sub(1, 1); o != 2 {
		t.Errorf("Sub(1,1) != %d", o)
	}
}
func TestSub1(t *testing.T) {
	if o := Sub1(1, 1); o != 2 {
		t.Errorf("Sub(1,1) != %d", o)
	}
}
