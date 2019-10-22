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
	l := len(list)
	if l <= 1 {
		return list
	}

	for i := 0; i < l; i++ {
		for j := 1 + i ; j < l-i ; j++ {
			if list[j] < list[i] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
	return list
}
