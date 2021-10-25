package main

import "fmt"
func work(x int) int {
	x++
	return x
}
func main() {
	m := make(map[int]int)
	for i:=0; i<10; i++{
		var s int
		fmt.Scan(&s)
		if result,ok := m[s]; ok {
			fmt.Print(result," ")
		} else {m[s] = work(s)
			fmt.Print(m[s]," ")
		}
	}


}
/* Если есть расчитанные значения для определенного числа,
	то не считать их заново,а взять готовые
*/