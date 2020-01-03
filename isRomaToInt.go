package main

import "fmt"

func main(){
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	dict := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	if s == "" {
		return 0
	}
	re := 0
	le := len(s)
	pre := 0
	next := 0
	for i:=0;i<le-1;i++{
		pre = dict[string(s[i])]
		next = dict[string(s[i+1])]
		fmt.Println(next)
		if pre < next {
			re = re + (next-pre)
		} else {
			re = re + pre + next
		}
		i++
	}
	return re
}
