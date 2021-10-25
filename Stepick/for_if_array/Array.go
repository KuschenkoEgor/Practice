package main

import "fmt"

func main()  {
	var workArray [4]uint64
	for id:=0; id< len(workArray); id++{
		var a uint64
		fmt.Scan(&a)
		workArray[id] = a
	}
	fmt.Scan(&workArray)
	var (
		para1/*,para2,para3*/ [2]uint64
		//swap uint64
	)
		for i:=0; i< len(para1); i++{
			fmt.Println("Введите пару №1")
			var p1 uint64
			fmt.Scan(&p1)
			para1[i] = p1
		}
		workArray[para1[0]],workArray[para1[1]] = workArray[para1[1]],workArray[para1[0]]
		/*		swap = workArray[para1[0]]
				workArray[para1[0]] = workArray[para1[1]]
				workArray[para1[1]] = swap

		for j:=0; j< len(para2); j++{
			fmt.Println("Введите пару №2")
			var p2 uint64
			fmt.Scan(&p2)
			para2[j] = p2
	}
				swap = workArray[para2[0]]
				workArray[para2[0]] = workArray[para2[1]]
				workArray[para2[1]] = swap

		for k:=0; k< len(para3); k++{
			fmt.Println("Введите пару №3")
			var p3 uint64
			fmt.Scan(&p3)
			para3[k] = p3
	}
				swap = workArray[para3[0]]
				workArray[para3[0]] = workArray[para3[1]]
				workArray[para3[1]] = swap*/

		fmt.Println(para1)
		/*fmt.Println(para2)
		fmt.Println(para3)*/
		for i:=0; i<len(workArray); i++{
			fmt.Print(workArray[i], " ")
		}
}
