package main

import "fmt"

var a string
var c = make(chan int)

func fun() {
	c <- 0
	a = "hello world"
}

func main() {
	go fun()
	<-c
	fmt.Println(a)

}
