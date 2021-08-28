package main

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// length converts number received as arguments or from STD output and prints the conversion from
// feet to meters and vice versa
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			r := bufio.NewReader(os.Stdin)
			text, _ := r.ReadString('\n')
			t, err = strconv.ParseFloat(text, 64)
			if err != nil {
				log.Fatalln("what are you doing")
			}
		}

		p := KGToPounds(t)
		kg := PoundsToKG(t)
		fmt.Printf("%v kg = %v pounds, %v  pounds = %v kg\n",
			t, p, t, kg)
	}
}

func PoundsToKG(u float64) float64 {
	return u * 0.453592
}

func KGToPounds(u float64) float64 {
	return u * 2.20462
}
