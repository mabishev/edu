package main

import "fmt"

func main() {
	x1 := 4
	x2 := 2
	test(&x1, &x2)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}
