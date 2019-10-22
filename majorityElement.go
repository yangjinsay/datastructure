package main

import "fmt"

func main(){
	nums := []int{1,2,3,3,4,3,3}
	result := majorityElement(nums)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	list := map[int]int{}
	l := len(nums)
	for _, v := range nums {
		if j, ok := list[v]; ok {
			list[v] = j + 1
		} else {
		list[v] = 1
		}
	}
	for num, n := range list {
		if n > l/2 {
		return num
		}
	}
	return 0
}
