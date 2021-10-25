package main

import (
	"fmt"
	"unicode/utf8"
)

func main () {
	var x , s string
	fmt.Scan(&x, &s)
	tx1 := []rune(x)
	tx2:= []rune(s)
	N := 0
	for i:=0; i<utf8.RuneCountInString(x)-1; i++ {
		if tx1[i] == tx2[0] && tx1[i+1] == tx2[1] {
			N = i
			break
		} else {N = -1}
 	}
		fmt.Println(N)
}

/*
func main() {
	var x, s string
	fmt.Scan(&x, &s)
	fmt.Println(strings.Index(x, s))
} */