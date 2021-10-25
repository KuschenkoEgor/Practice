package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scan(&a)
	rn := []rune(a)
	//rn1:= []rune(a)
	for i:=0; i<utf8.RuneCountInString(a); i++ {
		fmt.Print((rn[i]-48)*(rn[i]-48))
	}
}
