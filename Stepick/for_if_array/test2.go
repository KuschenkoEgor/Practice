package main

import "fmt"

func main() {
	var x, p, y float64
	fmt.Scan(&x, &p, &y)
	summ := x
	var i int
	for i = 1; ; i++ {
		summ = summ + (p/100)*summ
		if summ >= y {
			break
		}
		fmt.Println(summ)

	}
	fmt.Println(i)

}