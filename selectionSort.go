package main

import (
	"fmt"
)

/**
* 选择排序每次从未排序区间找出最小的元素，将其放在已排区间的末尾.
*/
func main() {
	list := []int{3, 5, 4, 1, 2, 6}
	fmt.Println("Before:", list)
	reList := bubbleSort(list)
	fmt.Println("After",reList)
}

func bubbleSort(list []int) []int {
	le := len(list)
	if le <= 1 {
		return list
	}
	minVal := list[0]
	minStart := 0
	for{
		for i := minStart; i < le; i++ {
			if minVal > list[i] {
				list[minStart],list[i] = list[i],list[minStart]
				minS
			}
			fmt.Println(list)
		}
		if minStart == le-1 {
			break
		}
		minStart += 1
	}
	return list
}
