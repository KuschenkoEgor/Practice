package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scan(&a)
	t1 := []rune(a)
	for i:=0; i<utf8.RuneCountInString(a); i++ {
		if strings.Count(a, string(t1[i])) >= 2 {
			continue
		} else {fmt.Print(string(t1[i]))}
	}
}
/* for _, ch := range a {
    if strings.Count(a, string(ch)) == 1 {
        fmt.Print(string(ch))
    }
}
*/