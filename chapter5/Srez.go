package main

import "fmt"

func main() {
   /* var x [5] float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83
    */
  array := [5] float64{ 1, 2, 3, 4, 5 }
	//x[] - srez
	x := array[1:4]
	fmt.Println("Array", array )
	fmt.Println("Srez", x )
}