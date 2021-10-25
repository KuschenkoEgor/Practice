package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scan(&a)
	r := []rune(a)
	N := 0
	if utf8.RuneCountInString(a) > 5 {
		for i := 0; i < utf8.RuneCountInString(a); i++ {
			if unicode.Is(unicode.Latin, r[i]) || unicode.IsDigit(r[i]) {
				N++
			}
		}

	}
	if N == utf8.RuneCountInString(a) {
		fmt.Print("Ok")
	} else {fmt.Print("Wrong password")}
}