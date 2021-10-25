package main

import "fmt"


func main()  {
	var N,i int
	fmt.Scan(&N)
	sl:= make([]int, N, 2*N)
	sl[0] = 0
	sl[1] = 1
	for i= 2; sl[i] < N ; i++ {
		sl[i] = sl[i-1] + sl[i-2]
		fmt.Print(sl[i], " ")
	}
	if sl[i] == N {
		fmt.Print(i)
	} else {fmt.Print("-1")}


}