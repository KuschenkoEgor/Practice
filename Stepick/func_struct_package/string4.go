package main

import (
	"fmt"
	"unicode/utf8"
)

func main (){

	var a string
	fmt.Scan(&a)
	tx1 := []rune(a)
	for i:=0; i< utf8.RuneCountInString(a); i++{
		if i%2 != 0 {
			fmt.Print(string(tx1[i]))
		}
	}
}

