package main

import ("fmt"
	"strings"
)// пакет используется для проверки ответа, не удаляйте его
type Stringer interface {
	String() string
}

type Battery struct {
	Charge string
}
func (b Battery) String() string  {
	N := strings.Count(b.Charge,"1")
	//b.Charge = strings.Repeat("X", N)
	Newstr := strings.Repeat("X", N)
	return fmt.Sprintf("[%10s] ", Newstr)
}

func main() {
	var text string
	fmt.Scan(&text)
	batteryForTest := Battery{Charge: text}
}