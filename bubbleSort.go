package main

import (
	"fmt"
)

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

	for i:=0;i<le;i++{
		for j:= i+1; j< le;j++{
			if list[j] < list[i]{
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
