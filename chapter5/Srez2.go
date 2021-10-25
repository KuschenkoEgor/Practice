package main

import "fmt"

func main() {
    slice1 := []int{1,2,3,4,5,6,76,7}
    slice2 := make([]int, 5)
    slice3 := []int{11, 24, 33, 11}
    copy(slice2, slice3) // [11,24,33,11,0]
    fmt.Println(slice1, slice2, slice3)
}