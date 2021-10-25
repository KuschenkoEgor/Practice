package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	text, _:= bufio.NewReader(os.Stdin).ReadString('\n')
	if strings.ToUpper(text[:2]) == text[:2] && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {fmt.Println("Wrong") }
}