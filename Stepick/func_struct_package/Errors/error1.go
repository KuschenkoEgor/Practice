package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	er := errors.New("ошибка")
	_,err := fmt.Scan(&a, &b)
	if err == nil {
		return a/b
	} else {return er}
}