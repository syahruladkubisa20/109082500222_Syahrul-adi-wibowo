package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	x, y, r int
}

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.x, &c1.y, &c1.r)
	fmt.Scan(&c2.x, &c2.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	d1 := jarak(c1.x, c1.y, p.x, p.y) <= float64(c1.r)
	d2 := jarak(c2.x, c2.y, p.x, p.y) <= float64(c2.r)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
