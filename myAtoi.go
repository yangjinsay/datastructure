package main

import (
	"strings"
	"math"
	"fmt"
)

func main(){
}

func myAtoi(str string) int {
}

func clean(s string) (sign int, abs string){
	//去除首位空格
	s = strings.TrimSpace(s)
	switch s[0]{
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}
}
