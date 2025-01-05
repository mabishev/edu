package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var sum int
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	fmt.Println(sum)
}
