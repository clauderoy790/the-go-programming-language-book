package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	nb := "123456789.3456576"
	res := comma(nb)
	res2 := nonRecursiveComma(nb)
	fmt.Println("comma: ", res)
	fmt.Println("non recursive: ", res2)
	strs := []stringComp{
		{"listen", "triangle"},
		{"listen", "silent"},
		{"listen", "integral"},
		{"oelhl", "hello"},
		{"salut", "hello"},
	}
	for _, str := range strs {
		fmt.Printf("%v is anagram of %v? %v\n", str.s1, str.s2, anagram(str.s1, str.s2))
		fmt.Printf("%v is anagram2 of %v? %v\n", str.s1, str.s2, anagram2(str.s1, str.s2))
		fmt.Printf("%v is anagram3 of %v? %v\n", str.s1, str.s2, anagram3(str.s1, str.s2))
	}
}

type stringComp struct {
	s1, s2 string
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
	m := make([]map[rune]int, 2)
	for i, s := range []string{s1, s2} {
		m[i] = make(map[rune]int)
		for _, r := range s {
			_, ok := m[i][r]
			if !ok {
				m[i][r] = 0
			}
			m[i][r]++
		}
	}
	return reflect.DeepEqual(m[0], m[1])
}

func anagram2(s1, s2 string) bool {
	slices := make([][]string, 2)
	for ind, str := range []string{s1, s2} {
		slices[ind] = strings.Split(str, "")
		sort.Slice(slices[ind], func(i, j int) bool {
			return slices[ind][i] < slices[ind][j]
		})
	}
	return strings.Join(slices[0], "") == strings.Join(slices[1], "")
}

func anagram3(s1, s2 string) bool {
	slice := []byRune{byRune(s1), byRune(s2)}
	for _, r := range slice {
		sort.Sort(r)
	}
	return string(slice[0]) == string(slice[1])
}

type byRune []rune

func (r byRune) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r byRune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRune) Len() int {
	return len(r)
}
