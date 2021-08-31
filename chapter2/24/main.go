package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println("the val: ", byte(i&1))
		fmt.Println("init called")
	}
}

func main() {
	nb := uint64(1 << 4)
	fmt.Println("bits count: ", nb)
}

func PopCount(x uint64) int {
	val := 0
	for i := 0; i < 8; i++ {
		val += int(pc[byte(x>>(i*8))])
	}
	return val
}

func PopCount23(x uint64) int {
	val := 0
	for i := 0; i < 64; i++ {
		if (1<<i)&x != 0 {
			val++
		}
	}
	return val
}

func PopCount24(x uint64) int {
	val := 0
	for x > 0 {
		val = val & (val - 1)
		val++
	}
	return val
}
