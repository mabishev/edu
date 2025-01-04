package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)

	for i := 0; i < len(a); i++ {
		a[i] = append(a[i], 5)
		fmt.Println(a)
	}

}
