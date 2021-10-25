package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var text string
	fmt.Scan(&text)
	tx1:= []rune(text)
	tx2:= []rune(text)
	tx2[0] = tx1[utf8.RuneCountInString(text)-1]
	for i:= 1; i<utf8.RuneCountInString(text); i++{
		tx2[i] = tx1[utf8.RuneCountInString(text)-i-1]
	}
	fmt.Println(tx1, tx2)
	N:=0
	for k:=0; k< utf8.RuneCountInString(text); k++{
		if tx1[k] == tx2[k] {
			N++
		}
	}
	if N == utf8.RuneCountInString(text) {
		fmt.Println("Палиндром")
	} else {fmt.Println("Нет")}
}

