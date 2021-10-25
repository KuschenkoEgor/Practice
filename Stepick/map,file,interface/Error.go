package main

import (
	"fmt"
	"unicode"
)

type customError uint

func (c customError) Error() string {
	return fmt.Sprintf("digit!!!: index %d", c)
}

func errorInString(str string) error {
	// Полезная работа со строкой проигнорирована
	for i, s := range str {
		if unicode.IsDigit(s) {
			return customError(i)
		}
	}
	return nil
}

func ExampleMyError() {
	err := errorInString("Строка1Строка")
	if err != nil {
		fmt.Printf("Ошибка обработана: %v\n", err)
	}
	if idx, ok := err.(customError); ok {
		fmt.Printf("Контекст: %d\n", idx)
	}

	// Output:
	// Ошибка обработана: digit!!!: index 6
	// Контекст: 6
}