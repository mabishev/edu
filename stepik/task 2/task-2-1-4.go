package main

import "fmt"

func main() {
	a := fibonacci(4)
	fmt.Println(a)
}
func fibonacci(n int) int {
	var j int
	var x, y int = 1, 1
	for i := 2; i < n; i++ {
		j = x + y
		y = x
		x = j
	}
	if n > 2 {
		return j
	}
	return 1

	/*
			func fibonacci(n int) int {
		  if n<=2 {
		     return 1
		  }
		  return fibonacci(n-2)+fibonacci(n-1)
		}
	*/
}
