package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var numbers []int
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}

	var count int
	for _, n := range numbers {
		if n > 0 {
			count++
		}
	}
	fmt.Println(count)
}
