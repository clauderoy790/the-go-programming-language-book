package calculator
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	val := 0
	for i :=0; i < 8;i++ {
		val += int(pc[byte(x>>(i*8))])
	}
	return val
}