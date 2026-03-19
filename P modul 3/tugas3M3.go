package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func dalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var a, b, c, d, e, f float64
	var x, y float64

	fmt.Scan(&a, &b, &c)
	fmt.Scan(&d, &e, &f)
	fmt.Scan(&x, &y)
	bulat := dalam(a, b, c, x, y)
	bulat2 := dalam(d, e, f, x, y)

	if bulat && bulat2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if bulat {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if bulat2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
