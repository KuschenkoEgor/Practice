package main

import "fmt"

type Document struct {
	Date   string
	Number string
	Page   int
}
type Card struct {
	Date string
	Name string
	Age  int
}

type PrintInterface interface {
	CheckData()
	//CheckDoc()
	//CheckCard()
}

func (d *Document) CheckData() {
	if d.Date != "" {
		fmt.Println("Date in the doc - correct")
	} else {
		fmt.Println("Not correct Date in doc")
	}
}
func (c *Card) CheckData() {
	if c.Date != "" {
		fmt.Println("Date in the Card - correct")
	} else {
		fmt.Println("Not correct Date in the Card")
	}
}


/*func PrintAnything(arr []PrintInterface) {
	for _, v := range arr {
		fmt.Println(v)
	}
}*/

func main() {
	doc := new(Document)
	card := new(Card)

	doc.Date = "01.22.11"
	doc.Number = "5"
	doc.Page = 11

	card.Date = "01.11.22"
	card.Age = 20
	card.Name = "Zhorik"
	//doc.CheckData()
	//card.CheckData()
	s1 := []PrintInterface{doc, card}
	for _, v := range s1 {
		v.CheckData()

	}
}
