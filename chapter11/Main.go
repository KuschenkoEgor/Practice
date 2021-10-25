package main

import "fmt"
import m "C:\Program Files\Go\src\goolang-book\chapter11>"

func main() {
    xs := []float64{1,2,3,4}
    avg := m.Average(xs)
    fmt.Println(avg)
}