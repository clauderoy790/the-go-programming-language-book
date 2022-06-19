package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

func main() {
	h1, h2 := "salut", "sal5t"
	fmt.Printf("different bits: %d\n", countDiffBits(h1, h2))
}

func countDiffBits(h1, h2 string) int {
	hash1 := sha256.New()
	hash1.Write([]byte(h1))
	final1 := hash1.Sum(nil)

	hash2 := sha256.New()
	hash2.Write([]byte(h2))
	final2 := hash2.Sum(nil)

	l := 0
	l1, l2 := len(final1), len(final2)
	switch {
	case l1 < l2:
		l = l1
	default:
		l = l2
	}

	diff := int(math.Abs(float64(l1 - l2)))

	for i := 0; i < l; i++ {
		if final1[i] != final2[i] {
			diff++
		}
	}

	return diff
}
