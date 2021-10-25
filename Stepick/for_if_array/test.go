package main

import (
	"fmt"
)

func main() {
	//var n int
	fmt.Println("Введите количество значений:")
	var n int
	fmt.Scan(&n)
	sum:=0
	for i:=1; i<=n; i++ {
		var k int
		fmt.Scan(&k)
		if k%8 == 0 && 99>k && k>9 {
			sum += k
		}
	}

		fmt.Println(sum)
}