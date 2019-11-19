package math

import "math"

func Digits(n int) int {
	return int(math.Log10(float64(n))) + 1
}
