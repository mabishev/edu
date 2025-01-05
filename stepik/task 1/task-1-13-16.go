package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	array := []int{}

	for a > 0 {
		x := a % 10
		if x != b {
			array = append(array, x)
			a /= 10
		} else {
			a /= 10
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i])

	}
}
