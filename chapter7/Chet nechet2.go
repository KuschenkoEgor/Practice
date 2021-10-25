package main

import "fmt"

func main() {
   x:= []int{11, 23, 43, 10, 33}
   for _,v := range x {
	pop := v/2
	if v % 2 == 0 {
		fmt.Println(pop, "True")
	} else {
		fmt.Println(pop, "False")
	}
   }
}
   
