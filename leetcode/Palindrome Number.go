package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	pol := 0
	digit :=0
	x1:=x
	for x1 > 0 {
		digit = x1%10
		x1 = x1/10
		pol = pol*10
		pol = digit + pol
	}

	if pol == x && x >= 0 {
		fmt.Print("true")
		return true
	} else {
		fmt.Print("false")
		return false}
}

func main() {
	isPalindrome(-121)
}