package main

import "fmt"


	/* func f() {
	 fmt.Println(x)
	}
	 func main() {
	 x := 5
	 f()
	} */ // Ne srabotaet

func f(x int) {
    fmt.Println(x)
}
func main() {
    y := 5
    f(y)
}