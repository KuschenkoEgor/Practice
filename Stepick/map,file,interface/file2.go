package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Print("error")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println(reader.ReadString(';'))

}