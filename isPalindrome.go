package main

import (
	"fmt"
)

func main(){
	r := isPalindrome(111)
	fmt.Println(r)
}
func isPalindrome(x int) bool {

	inPut := x
	if x < 0 {
	return false
	}

	if x >= 0 && x < 10 {
		return true
	}

	list := []int{}
	offSet := 1
	mod := 0
	for {
		mod = x%10
		x = x/10
		if mod == 0 && x == 0 {
			break
		}
		if x != 0 {
			offSet = offSet* 10
		}
		list = append(list, mod)
	}

	// splice
	splice := 0
	for _,v := range list {
		splice = splice + v * offSet
		offSet = offSet/10
	}
	if splice != inPut {
		return false
	}
	return true
}
