package main

import "fmt"

func first() {
    fmt.Println("1st")
}
func second() {
    defer  third()
    fmt.Println("2nd")
}
func third() {
    fmt.Println("3nd")
}
func main() {
    //defer  first()  otkladivaet v konets
    second()
    first()
    
}