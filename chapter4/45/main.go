package main

import "fmt"

func main() {
	sl := []string{"1", "1", "2", "2", "2"}
	sl = eliminateAdjacent(sl)
	fmt.Println("eliminated adjacents: ", sl)
}

func eliminateAdjacent(sl []string) []string {
	var final []string
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] != sl[i+1] {
			final = append(final, sl[i])
		}
	}
	return final
}
