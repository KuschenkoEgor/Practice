package main

import (
	"fmt"
	"math"
)

func treug (a, b float64 ) float64 {
	return math.Sqrt(a*a + b*b)
}
func main() {
	var c, d float64
	fmt.Scan(&c, &d)
	fmt.Println(treug(c,d))
}