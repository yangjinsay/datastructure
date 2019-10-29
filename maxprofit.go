package main

import (
	"fmt"
)

func main() {
	list := []int{7, 1, 5, 3, 6, 4}
	re := maxProfit(list)
	fmt.Println(re)
}

func maxProfit(prices []int) int {
	pre := 0
	value := 0
	for i, li := range prices {
		if i == 0 {
			pre = li
			continue
		}
		if li > pre {
			value = value + (li - pre)
		}
		pre = li
	}
	return value
}
