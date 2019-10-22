package main

import "fmt"

func main(){
	result := pow(2,-2)
	fmt.Println(result)
}

func pow(x int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 1/pow(x,-n)
	}

	if n%2 == 0 {
		return  pow(x * x, n/2)
	}
	return x * pow(x, n-1)
}
