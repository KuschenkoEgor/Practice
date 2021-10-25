package main

import "fmt"


func max(args ...int)int {
     maxEl := 0
     for _, value := range args{ // value = x[i]
	if value > maxEl {
		maxEl = value
		}
     }
      return maxEl
}
func print() {
fmt.Println(max(11, 33, 32, 45, 12))
}

func main() {
    print()
}