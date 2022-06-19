package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = rotate(arr, 2)
	fmt.Println("rotated: ", arr)
}

func rotate(arr []int, i int) []int {
	l := len(arr)
	s := arr[i-1 : l-i*2]
	s1 := arr[l-1-i:]
	s2 := arr[:i]
	return append(s, append(s1, s2...)...)
}
