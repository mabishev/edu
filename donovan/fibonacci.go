package edu

import (
	"fmt"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
func Fib() {
	fmt.Println(fib(5))

}
