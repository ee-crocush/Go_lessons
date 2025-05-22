package tasks

import "math"

// CountLetters подсчитывает количество символов в строке.
func CountLetters(s string, l rune) int {
	c := 0
	for _, b := range []rune(s) {
		if b == l {
			c++
		}
	}
	return c
}

func SomeMath(rad float64) float64 {
	return math.Sin(rad) * math.Cos(rad)
}
