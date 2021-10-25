package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x > math.MaxInt32 -1 || x < -math.MaxInt32  {
		return 0

	}
	digit :=0
	Newx:=0
	for x>0 || x<0{
		digit = x%10
		x = x/10
		Newx = Newx*10
		Newx = Newx + digit
	}
	if Newx > math.MaxInt32 -1 || Newx < -math.MaxInt32 {
		return 0
	} else {
		fmt.Print(Newx)
		return Newx}
}

func main(){
	fmt.Print(reverse(-123))
	fmt.Println(1534236469-2147483646 )
}