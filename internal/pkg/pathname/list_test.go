package pathname

import (
	"testing"
)

func TestList_Get(t *testing.T) {
	list := NewList(
		New("images/001.png"),
		New("images/002.png"),
		New("images/003.png"),
		New("images/004.png"),
		New("images/005.png"),
	)

	actual := list.Get(2)
	expected := New("images/003.png")

	if !actual.Equals(expected) {
		t.Errorf("Actual: %v, Expected: %v", actual, expected)
	}
}

func TestList_Equals(t *testing.T) {
	list1 := NewList(
		New("images/001.png"),
		New("images/002.png"),
		New("images/003.png"),
		New("images/004.png"),
		New("images/005.png"),
	)
	list2 := NewList(
		New("images/001.png"),
		New("images/002.png"),
		New("images/003.png"),
		New("images/004.png"),
		New("images/005.png"),
	)

	if !list1.Equals(list2) {
		t.Fail()
	}
}

func TestList_Find(t *testing.T) {
	list := NewList(
		New("images/001.png"),
		New("images/002.png"),
		New("images/003.png"),
		New("images/004.png"),
		New("images/005.png"),
	)

	actual := list.Find(&Pathname{"images/001.png"})
	expected := &Pathname{"images/001.png"}

	if !actual.Equals(expected) {
		t.Errorf("Actual: %v, Expected: %v", actual, expected)
	}
}

func TestList_SortStable(t *testing.T) {
	actual := NewList(
		New("images/002.png"),
		New("images/005.png"),
		New("images/004.png"),
		New("images/003.png"),
		New("images/001.png"),
	)

	actual.SortStable()

	expected := NewList(
		New("images/001.png"),
		New("images/002.png"),
		New("images/003.png"),
		New("images/004.png"),
		New("images/005.png"),
	)

	if !actual.Equals(expected) {
		t.Errorf("Actual: %v, Expected: %v", actual, expected)
	}
}
