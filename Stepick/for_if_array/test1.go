package main

import "fmt"

func main() {
	var a, n int
	max:=0
	for fmt.Scan(&a); a >0; fmt.Scan(&a) {
		//var a int
		//fmt.Scan(&a)
		//n := 0
		if a > max {
			max = a
			n = 1
			//fmt.Println("Проверка в if", n)
		} else if a == max{
			n++
			//fmt.Println("Проверка else if", n)
		}
		var _ int
	}
	fmt.Println("n=", n)
	fmt.Println("max=", max)
}