package main

import "fmt"

func average(xs []float64) float64 {
     total := 0.0
    for _, v := range xs {
        total += v
}
	return total / float64(len(xs))
}


func main() {
    xxx := []float64{98,93,77,82,83} // Ne obyazatelno ispolzovat XS kak v average
    fmt.Println(average(xxx))
}

