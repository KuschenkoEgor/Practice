package main

import "fmt"

func main() {
	var a int = 2
	//a := 2
	b := &a
	a = 10

	*b = 15
	fmt.Println(a)
	fmt.Println(b)
	c := &a
	*c = 20
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)
	//var d *int
	d := new(int)
	fmt.Println(d, "___")
	*d = 12
	*c = *d
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("---")
	*d = 13
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
	c = d
	*c = 14
	fmt.Println(a, "2")
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
