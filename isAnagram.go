package main

import "fmt"



func isAnagram(s string, t string) bool {
	if s == t || sort(s) == sort(t) {
		return true
	} else {
		return false
	}
}

funct sort(s sting) string {
	le := len(s)
	if le == 1 {
		return s
	}
	strList := []rune(s)
	for i:=0;i<le-1;i++{
		for j:= i+1;j<le;j++{
			if strList[j] > strList[i] {
				strList[i], strList[j] = strList[j],strList[i]
			}
		}
	}
	return string(strList)
}
