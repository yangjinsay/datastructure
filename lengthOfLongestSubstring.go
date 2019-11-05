package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
}

func lengthOfLongestSubstring(s string) int {
	if isRepeat(s) {
		return len(s)
	}

	subStr := ""
	maxlen := 1
	curlen := 0
	le := len(s)
	for j := 0; j < le-1; j++ {
		if maxlen > len(s[j:len(s)]){
			return maxlen
		}
		for i := j + 1; i < le; i++ {
			tmpStr := s[j:i+1]
			if isRepeat(tmpStr) {
				subStr = tmpStr
				curlen = len(subStr)
				if maxlen < curlen {
					maxlen = curlen
				}
			} else {
				break
			}
		}
	}
	return maxlen
}

func isRepeat(str string) bool {
	if len(str) == 1 {
		return true
	}
	for _, v := range str {
		newStr := strings.Replace(str, string(v), "", -1)
		if len(str)-1 > len(newStr) {
			return false
		}
	}
	return true
}
