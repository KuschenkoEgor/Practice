package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//const now = 1589570165
	//nowtime := time.Unix(now,0)
	a := bufio.NewReader(os.Stdin)
	timestr, _ := a.ReadString('\n')
	timestr = strings.TrimSuffix(timestr,"\n")
	dur, _ := time.ParseDuration(timestr)
	fmt.Println(dur)
}
