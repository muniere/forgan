package pathname

import (
	"testing"
)

func TestPathname_Equals(t *testing.T) {
	e1 := New("images/001.png")
	e2 := New("images/001.png")

	if !e1.Equals(e2) {
		t.Fail()
	}
}
