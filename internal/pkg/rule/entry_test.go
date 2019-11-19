package rule

import (
	"testing"
)

func TestRule_Equals(t *testing.T) {
	e1 := New("images/001.png", "images/011.png")
	e2 := New("images/001.png", "images/011.png")

	if !e1.Equals(e2) {
		t.Fail()
	}
}
