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
	for i=0;i<le;i++{
		pre = dict[string(dict[s[i]])
		if i < (le -1) {
			next = dict[string(dict[s[i+1]])
		}
		if pre < next {
			re +=(next-pre)
			i++
		} else {
			re += pre
		}
	}
	return re
}
