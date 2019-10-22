package main

import "fmt"

func main(){
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if result != nil {
		fmt.Println(result)
	}
}

func twoSum(nums []int, target int) []int {
	l := len(nums)
	list := map[int]int{}
	if l <= 1 {
		return nil
	}
	for i, v := range nums {
		if j,ok := list[target - v]; ok {
			return []int{i, j}
		}
		list[v] = i
	}
	return nil
}
