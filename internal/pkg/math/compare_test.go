package math

import (
	"testing"
)

func TestMath_Max(t *testing.T) {
	t.Run("Different Values", func(t *testing.T) {
		actual := Max(3, 5)
		expected := 5

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})

	t.Run("Same Values", func(t *testing.T) {
		actual := Max(5, 5)
		expected := 5

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})
}

func TestMath_Min(t *testing.T) {
	t.Run("Different Values", func(t *testing.T) {
		actual := Min(3, 5)
		expected := 3

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})
	t.Run("Same Values", func(t *testing.T) {
		actual := Min(5, 5)
		expected := 5

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})
}
