package rule

import (
	"testing"

	"github.com/muniere/forgan/internal/pkg/pathname"
)

func TestList_Get(t *testing.T) {
	list := NewList(
		New("images/001.png", "images/011.png"),
		New("images/002.png", "images/012.png"),
		New("images/003.png", "images/013.png"),
	)

	actual := list.Get(1)
	expected := New("images/002.png", "images/012.png")

	if !actual.Equals(expected) {
		t.Errorf("Actual: %v, Expected: %v", actual, expected)
	}
}

func TestList_Equals(t *testing.T) {
	list1 := NewList(
		New("images/001.png", "images/011.png"),
		New("images/002.png", "images/012.png"),
		New("images/003.png", "images/013.png"),
	)
	list2 := NewList(
		New("images/001.png", "images/011.png"),
		New("images/002.png", "images/012.png"),
		New("images/003.png", "images/013.png"),
	)

	if !list1.Equals(list2) {
		t.Fail()
	}
}

func TestList_Find(t *testing.T) {
	list := NewList(
		New("images/001.png", "images/011.png"),
		New("images/002.png", "images/012.png"),
		New("images/003.png", "images/013.png"),
	)

	actual := list.FindByDst(pathname.New("images/011.png"))
	expected := New("images/001.png", "images/011.png")

	if !actual.Equals(expected) {
		t.Errorf("Actual: %v, Expected: %v", actual, expected)
	}
}
