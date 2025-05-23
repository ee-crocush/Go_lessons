package main

import (
	"fmt"
	"unicode"
)

func main() {
	cities := []string{
		"Moscow",
		"Washington",
		"New-York",
		"Kiev",
		"Vitebsk",
		"Kishinev",
		"Vladivostok",
		"Ekaterinburg",
		"Bratsk",
	}

	fmt.Println(cities[0])

	for i := 1; i < len(cities); i++ {
		p, n := cities[i-1], cities[i]
		if !isCorrectCityName(p, n) {
			fmt.Println(cities[i], "-", "WRONG!")
			break
		}

		fmt.Println(cities[i], "-", "OK!")
	}
}

func isCorrectCityName(prev, next string) bool {
	if len(prev) == 0 || len(next) == 0 {
		return false
	}

	lastLetterPrev := rune(prev[len(prev)-1])
	firstLetterNext := rune(next[0])

	if unicode.ToLower(lastLetterPrev) == unicode.ToLower(firstLetterNext) {
		return true
	}

	return false
}
