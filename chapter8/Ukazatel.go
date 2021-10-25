package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 10
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}