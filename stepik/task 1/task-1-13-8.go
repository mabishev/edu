package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var n2 int
		fmt.Scan(&n2)
		if n2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
