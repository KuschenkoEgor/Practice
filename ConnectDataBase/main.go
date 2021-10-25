package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql","debian-sys-maint:jaguar228@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Подключено к MySQL")
}