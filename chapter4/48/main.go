package main

import (
	"fmt"
	"unicode"
)

func main() {
	counts := countChars("salut les ami, this is working 12341234332211")
	for k, v := range counts {
		fmt.Printf("k: %s, v: %+v\n", k, v)
	}
}

func countChars(str string) map[string]map[rune]int {
	counts := make(map[string]map[rune]int)

	for _, r := range str {
		if unicode.IsDigit(r) {
			addRuneToCategory(counts, "digit", r)
		}
		if unicode.IsLetter(r) {
			addRuneToCategory(counts, "letter", r)
		}
		if unicode.IsNumber(r) {
			addRuneToCategory(counts, "number", r)
		}
	}

	return counts
}

func addRuneToCategory(m map[string]map[rune]int, category string, r rune) {
	if m[category] == nil {
		m[category] = make(map[rune]int)
	}
	m[category][r]++
}
