package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	fmt.Println(str)
}
