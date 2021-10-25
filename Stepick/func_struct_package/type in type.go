package main

import "fmt"

type Person struct {
	Name string
}

func (n Person) Talk() {
	fmt.Println("Hello, I'm ", n.Name)
}

type Android struct {
	Person
	Model string
}

func (m Android) SerialNumb() {
	fmt.Println("My number is:" , m.Model)
}


func main()  {
	name := Person{"Ahsot"}
	md := Android{name, "228"}

	md.Person.Talk()
	md.Talk()
	md.SerialNumb()

}