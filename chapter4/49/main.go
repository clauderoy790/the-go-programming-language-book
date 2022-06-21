package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	p := point{1, 2}
	t := triangle{3}
	c := combined{p, t, "salut"}
	fmt.Println(c.point.x)

	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	m := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		if word == "stop" {
			break
		}
		m[word]++
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
	fmt.Println("THE COUNTS: ", m)
}

type point struct {
	x, y int
}

type triangle struct {
	edges int
}

type combined struct {
	point
	triangle
	name string
}

func get() t {
	v := t{}
	return v
}

type t struct {
	name string
}
