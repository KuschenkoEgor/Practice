package main

import "fmt"

/*func square(x *float64) {
	*x = 3
    *x = *x * *x
}
func main() {
    y := 1.5
    square(&y)
    fmt.Println(y)
}*/

func sq(x float64) float64{
	fmt.Println(&x) 
	fmt.Println(*&x) // * - znachenie adresa (Razimenovanie)
	x = x * x
	return x
}

func sqPointer(x *float64) {
	*x = *x * *x
	
	
}

func main() {
	var z float64 = 2
	z = sq(z)
	fmt.Println(z)
	sqPointer(&z) // Adres korobki "&"
	fmt.Println(z)
}