package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	a1 := strings.Split(a, "")
	fmt.Print(strings.Join(a1, "*"))
}
