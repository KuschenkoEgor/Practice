package main
import "fmt"

func main(){
	var N int
	fmt.Scan(&N)
	sl:=make([]int, N, 2*N)
	sum:=0
	for i:=0; i<N; i++{
		var a int
		fmt.Scan(&a)
		sl[i] = a
		if sl[i] > 0 {
			sum ++
		}
	}
	fmt.Println(sum)
}
