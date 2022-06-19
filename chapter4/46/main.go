package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Salut     les      amis comment   ca     va?"
	squashed := squash([]byte(str))
	fmt.Println("squased: ", squashed)
}

func squash(bytes []byte) string {
	spaceStart := -1
	consequentSpace := 0
	for i := 0; i < len(bytes); i++ {
		if unicode.IsSpace(rune(bytes[i])) {
			if spaceStart == -1 {
				spaceStart = i
			}
			consequentSpace++
		} else {
			if consequentSpace > 1 {
				bytes = append(bytes[:spaceStart+1], bytes[i:]...)
				i -= consequentSpace
			}
			consequentSpace = 0
			spaceStart = -1
		}
	}
	return string(bytes)
}
