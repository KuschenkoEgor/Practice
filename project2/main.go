package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	Age uint
	Money int
	Avg_grades float64
	Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string{
	return fmt.Sprintf("User name is: %s. He is %d and he " +
		"has money equal: %d",u.Name,u.Age,u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) { //Request отслеживает передачу, w - позволяет что-либо показывать пользователю
	bob:= User{"Bob", 25,-50,4.2,0.8, []string{"Football","Hockey"}}
	//fmt.Fprintf(w,`<h1> Main Text</h1>
		//<b>Main text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html") //Подгрузка HTML шаблона на гл страничке
	tmpl.Execute(w,bob) // Подгрузка шаблона на страницу и добавление данных

}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}
func HandleRequest() {
	http.HandleFunc("/", home_page) // Отслеживает переход по URL и вызывает метод
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) // создание сервера и отслеживание порта
}
func main() {
	HandleRequest() // Отслеживаем адреса и запускаем сервер

}

