package main

import "fmt"

type Rectangle struct {
	x,y float64
}

func (r *Rectangle) area() *Rectangle {
	r.x *= 2
	r.y += 2
	return r
}

func main()  {
	f := Rectangle{4, 8}
	f.area()
	fmt.Println(f)

}