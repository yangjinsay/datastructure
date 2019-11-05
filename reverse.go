package main

import "fmt"

func main(){
	splice := reverse(-1534236469)
	fmt.Println(splice)
}

func reverse(x int) int {

	if x/10 == 0 {
		return x
	}
	symbol := 1
	if x < 0 {
		symbol = -1
	}
	x = x * symbol
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
		if (splice < (-1<<31)) || (splice > (1<<31 -1)) {
			return 0
		}
		offSet = offSet/10
	}
	return splice * symbol
}
