package main

import "fmt"

func insert_sort() {
	A := [5]int
	for k,_ := range A {
		if (k>0 && A[k-1] > A[k]) {
			A[k],A[k-1] = A[k-1], A[k]
			k -= 1
		}
	}
}
func main()  {
	B := [5]int{11,22,12,2,3}
	fmt.Print(insert_sort(B))
}


