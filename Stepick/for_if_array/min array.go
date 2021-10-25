package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	sl := make([]int, N, 2*N)
	summ := 0
	var min int
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		sl[i] = a
	}
	min = sl[0]
	for i:=0; i<N; i++{
		if sl[i] < min {
			min = sl[i]
			summ = 1
		} else if sl[i] == min {
			summ ++
		}
	}
	fmt.Println(summ)
}