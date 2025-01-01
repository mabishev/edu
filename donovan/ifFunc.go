package edu

import "fmt"

func f(x int) int {
	return x * 2
}
func If() {
	if x := f(1); x == 4 {
		fmt.Println(x, "this if")
	} else {
		fmt.Println("0")
	}
}
