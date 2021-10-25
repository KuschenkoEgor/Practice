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
	S := strings.Split(timeStr, ",")
	firstTime, _ := time.Parse("02.01.2006 15:04:05", S[0])
	secondTime, _ := time.Parse("02.01.2006 15:04:05", S[1])
	if firstTime.After(secondTime) {
		fmt.Println(firstTime.Sub(secondTime))
	} else {fmt.Println(secondTime.Sub(firstTime)) }
}