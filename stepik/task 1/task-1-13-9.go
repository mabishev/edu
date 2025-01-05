package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	min := 9223372036854775807

	for i := 0; i < n; i++ {
		var cur int
		fmt.Scan(&cur)
		if cur == min {
			count++
		}
		if cur < min {
			min = cur
			count = 1
		}
	}
	fmt.Println(count)
}
