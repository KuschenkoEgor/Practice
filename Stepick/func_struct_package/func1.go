package main

import "fmt"

func sumInt(a ...int) (int,int) {
	summ := 0
	N := 0
	for _, elem:= range a {
		summ += elem
		N++
	}

	return N, summ
}
func main()  {
	a, b := sumInt(1,0)
	fmt.Print(a,b)
}