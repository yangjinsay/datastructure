package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(slice, 8))
}

func binarySearch(list []int, target int) int {
	var left, right int = 0, (len(list) - 1)
	for left <= right {
		middle := (left + right) / 2
		if target > list[middle] {
			left = middle + 1
		}
		if target < list[middle] {
			right = middle - 1
		}
		if target == list[middle] {
			return middle
		}
	}
	return -1
}
