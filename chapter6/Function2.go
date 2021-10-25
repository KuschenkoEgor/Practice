package main

import "fmt"

func f() (int, int) {
    return 5, 6
}

func main() {
    x, _ := f() // x, y := ...
    fmt.Println(x * 2)
}