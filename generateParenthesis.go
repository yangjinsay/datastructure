package main

import "fmt"

func main() {
	l := generateParenthesis(1)
	fmt.Println(l)
}

func generateParenthesis(n int) []string {
	generateOneByOne(0, 0, n, "")
	return result
}

var result []string

func generateOneByOne(left, right, n int, re string) {
	if left == n && right == n {
		result = append(result, re)
		return
	}

	if left < n {
		generateOneByOne(left+1, right, n, re+"(")
	}
	if left > right && right < n {
		generateOneByOne(left, right+1, n, re+")")
	}
}
