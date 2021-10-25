package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {

		var summ, st int
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			st, _ = strconv.Atoi(s.Text())
			summ = summ + st
		}
		io.WriteString(os.Stdout, strconv.Itoa(summ))
	}
