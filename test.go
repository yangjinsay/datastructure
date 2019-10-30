package main

import "fmt"

func main(){
	//a := []int{1,2,3}
	b := []int{8,9,10}
	d := 6
	re := append(d,b...)
	//re := append(a,d)
	//re = append(re,b...)
	fmt.Println(re)
}
