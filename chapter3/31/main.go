//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	f, err := os.Create("data.svg")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)

	_, err = writer.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	if err != nil {
		panic(err)
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			col := getColor(ay, by, cy, dy)
			_, err = writer.WriteString(fmt.Sprintf("<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				col, ax, ay, bx, by, cx, cy, dx, dy))
			if err != nil {
				panic(err)
			}
		}
	}
	_, _ = writer.WriteString("</svg>")
	writer.Flush()
}

func getColor(points ...float64) string {
	sum := .0
	for _, p := range points {
		sum += p
	}
	avg := sum / float64(len(points))
	min, max := float64(0x0000ff), float64(0xff0000)
	ratio := avg / float64(height)
	hex := min + (ratio)*(max-min)
	hexInt := int(hex)
	return fmt.Sprintf("#%X", hexInt)
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	ok := true
	for _, nb := range []float64{sx, sy} {
		if math.IsNaN(nb) || math.IsInf(nb, 0) {
			ok = false
			break
		}
	}
	return sx, sy, ok
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
