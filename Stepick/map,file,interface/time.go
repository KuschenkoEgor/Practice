package main

import (
	"fmt"
	"time"
)

func main(){
	var a string
	fmt.Scan(&a)

	firstTime, _ := time.Parse(time.RFC3339, a)
	fmt.Println(firstTime.Format(time.UnixDate))

}
