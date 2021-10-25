package main

import (
	"strconv"
	"unicode"
)

func adding(x,y string) int64 {
	rsx := []rune(x)
	rsy := []rune(y)
	var Newx string
	var Newy string
	for i := 0; i < len(x); i++ {
		if unicode.IsDigit(rsx[i]) {
			Newx = Newx + string(rsx[i])
		}
	}
	for i := 0; i < len(y); i++ {
		if unicode.IsDigit(rsy[i]) {
			Newy = Newy + string(rsy[i])
		}
	}
	resx,err := strconv.ParseInt(Newx,10,64)
	if err != nil {
		panic(err)
	}

	resy,err := strconv.ParseInt(Newy,10,64)
	if err != nil {
		panic(err)
	}
	return resx + resy
}