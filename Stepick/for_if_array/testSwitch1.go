package main

import "fmt"

func main() {
	v := 100
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		//fallthrough
	case 1:
		fmt.Println(1)

	default:
		fmt.Println("default")
	}
}