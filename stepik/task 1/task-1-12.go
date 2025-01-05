package main

import "fmt"

func main() {
	var N int
	var numbers []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	for i := 0; i < len(numbers); i += 2 {
		fmt.Printf("%v ", numbers[i])
	}
}
