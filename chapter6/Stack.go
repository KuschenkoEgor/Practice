package main

import "fmt"

func main() {
    fmt.Println(f1())
   
}
func f1() int {
    fmt.Println("f1")
    return f2()
 
}
func f2() int {
    fmt.Println("f2")
    return 3

}