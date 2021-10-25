package main

import "fmt"
// Zamikanie
func main() {
    x := 0
    increment := func() int {
        x+=2
        return x
    }
    fmt.Println(increment()) // 2 
    fmt.Println(increment()) 
    fmt.Println(increment())  	
}