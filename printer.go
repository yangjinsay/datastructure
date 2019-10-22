package main

import (
	"fmt"
	"time"
)

func main() {
	for {

		time.Sleep(time.Second * 1)
		go one()
		go two()
		go three()
	}
}

func one() {
	fmt.Println("One")
}

func two() {
	fmt.Println("Two")
}

func three() {
	fmt.Println("Three")
}
