package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	a := bufio.NewReader(os.Stdin)
	timeStr, _ := a.ReadString('\n')
	timeStr = strings.TrimSuffix(timeStr, "\n")
	firstTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

	if firstTime.Hour() > 13 {
		New:=firstTime.AddDate(0,0,1)
		fmt.Println(New.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
	}
}