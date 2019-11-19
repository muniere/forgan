package math

import (
	"testing"
)

func TestMath_Digits(t *testing.T) {
	t.Run("1-digit", func(t *testing.T) {
		actual := Digits(1)
		expected := 1

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})

	t.Run("2-digits", func(t *testing.T) {
		actual := Digits(10)
		expected := 2

		if actual != expected {
			t.Errorf("Actual: %v, Expected: %v", actual, expected)
		}
	})
}
