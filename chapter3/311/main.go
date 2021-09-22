package main

import "fmt"

func main() {
	r1, r2 := '3', 'g'

	fmt.Printf("str: %v, rune: %v,int32: %v, int32ToR: %v\n", string(r1), r1, int32(r1), rune(int32(r1)))
	fmt.Printf("str: %v, rune: %v,int32: %v, int32ToR: %v\n", string(r2), r2, int32(r2), rune(int32(r2)))
}
