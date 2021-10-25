package main
import "fmt"

func main(){
	var N int
	fmt.Scan(&N)
	sl:=make([]int, N, 2*N)
	for i:=0; i<N; i++{
		var a int
		fmt.Scan(&a)
		sl[i] = a
		if i%2 == 0 {
			fmt.Print(sl[i]," ")
		}
	}
}