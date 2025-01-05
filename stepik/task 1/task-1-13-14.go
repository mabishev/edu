package main

import "fmt"

func main() {
	var a, x, y int
	fmt.Scan(&a)
	y = 1
	count := 0
	for x < a {
		i := x + y
		y = x
		x = i
		count++
	}
	if x != a {
		fmt.Print(-1)
	} else {
		fmt.Println(count)
	}
}
