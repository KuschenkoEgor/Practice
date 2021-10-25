package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	max :=0
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
		if array[i] > max {
			max = array[i]
		}
	}
	fmt.Println(max)
}
