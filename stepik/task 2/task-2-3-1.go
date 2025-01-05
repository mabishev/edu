package main

import "fmt"

func main() {
	x1 := 2
	x2 := 2
	test(&x1, &x2)
}

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
