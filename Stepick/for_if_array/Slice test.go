package main

import "fmt"

func fnA(a [3]int) {
	a[1] = 15
}

func fnB(a []int) {
	a[1] = 15
}

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	a[2] = 22
	fnA(a)
	fnB(b)

	fmt.Println(a)
	fmt.Println(b)
}
