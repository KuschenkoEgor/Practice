package main

import (
	"fmt"
	"strconv"
)

fn := func(uint) uint {
	var a uint
	fmt.Scan(&a)
	stra := strconv.FormatUint(uint64(a),64)
	rs := []rune(stra)
	var NewA string
	for i:=0; i<len(rs);i++ {
		if rs[i]%2==0 && rs[i] != 0 {
			NewA = NewA + string(rs[i])
		}
	}
	resa,_ := strconv.ParseUint(NewA,10,64)
	if resa == 0 {
		return 100
	} return resa
}

func main()  {



}
