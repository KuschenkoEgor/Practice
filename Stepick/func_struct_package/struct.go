package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func Area(c *Circle) float64 {
	 s := (3.14 * c.r * c.r)
	 return s
}

func plus(p *Circle) *Circle  {
	p.x *= 2
	return p


}

func main()  {
	a := Circle{0,0,10}
	fmt.Println(Area(&a))
	p := Circle{2,0,10}
	plus(&p)
	fmt.Println(p)
}