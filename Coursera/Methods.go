package main

import "fmt"

type EDocument struct {
	Number        string
	Date          string
	NumberOfPages int
	FooterDoc     Footer
}

type Footer struct {
	Author string
}

func (doc *EDocument) ClearPages() {
	doc.NumberOfPages = 0
	fmt.Println(doc.NumberOfPages)
}
func (doc *EDocument) Rename() {
	doc.FooterDoc.Author = "Allah"
	fmt.Println("New name is:", doc.FooterDoc.Author)
}

func main() {
	doc1 := EDocument{
		Number:        "001",
		Date:          "19.08.21",
		NumberOfPages: 3,
		FooterDoc: Footer{
			Author: "Author-1",
		},
	}
	var doc2 EDocument
	doc2.Number = "002"
	doc2.Date = "19.08.22"
	doc2.NumberOfPages = 8
	doc2.FooterDoc.Author = "Zhora"

	doc1.ClearPages()
	doc2.Rename()
}
