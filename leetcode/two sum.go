package main

import "fmt"
func twoSum(nums []int, target int) []int {
	a := make([]int, 2, 4)
	for i:=0; i<len(nums); i++ {
		for k:=1; k<len(nums);k++ {
			if nums[i] + nums[k] == target && k!=i {
				a = []int{i,k}
				break
			}
		}
	}

	return a
}

func main(){
	a := []int{3, 3}
	fmt.Println(twoSum(a, 6))
}
