package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	nb := "123456789.3456576"
	res := comma(nb)
	res2 := nonRecursiveComma(nb)
	fmt.Println("comma: ", res)
	fmt.Println("non recursive: ", res2)
	strs := []stringComp{
		stringComp{"listen","triangle"},
		stringComp{"listen","silent"},
		stringComp{"listen","integral"},
		stringComp{"oelhl","hello"},
		stringComp{"salut","hello"},
	}	
	for _, str := range strs {
		fmt.Printf("%v is anagram of %v? %v\n",str.s1,str.s2,anagram(str.s1,str.s2))
	}
}

type stringComp struct {
	s1,s2 string
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	dotI := strings.Index(s, ".")
	if dotI != -1 {
		n = dotI
	}
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// comma inserts commas in a non-negative decimal integer string.
func nonRecursiveComma(s string) string {
	i := 0
	res := ""
	for len(s)-i > 3 {
		res += s[i:i+3] + ","
		i += 3
	}
	res += s[i:]
	return res
}

func anagram(s1, s2 string) bool {
	m := make([]map[rune]int,2)
	for i, s := range []string{s1,s2} {
		m[i] = make(map[rune]int)
		for _, r := range s {
			_, ok := m[i][r]
			if !ok {
				m[i][r] = 0
			}
			m[i][r]++
		}
	}
	return reflect.DeepEqual(m[0],m[1])
}