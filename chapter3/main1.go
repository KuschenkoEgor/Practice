package main

import "fmt"

func main() {
	var x string //= "Hello, mafaka". If const INSTEAD var - Can't change
	x =  "Hello, mafaka"
	fmt.Println(x)
	x = x + "228" // x += "228"
	fmt.Println(x)
	}