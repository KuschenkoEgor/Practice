package main

import "fmt"

func main() {
	fmt.Println(len("Hello, Nibba"))
	fmt.Println("Hello, Nibba"[1])
	fmt.Println("Hello" + "Nibba")
	fmt.Println((true && false) || (false && true) || !(false && false))
	}