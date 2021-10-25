package main

import "fmt"
import "math"


type Circle struct {
	x, y, r float64
	}

 
 func circleArea(c *Circle) float64 {
	//*c =Circle{0, 0, 1}
	return math.Pi * (c.r) * (c.r)
 }
 
 func main() {
 y := Circle{4, 1, 10}
 fmt.Println(circleArea(&y))
 }