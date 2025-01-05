package main

import "fmt"

func main() {
	var n, c, d uint
	fmt.Scan(&n, &c, &d)
	for c <= n && c%d != 0 {
		fmt.Println(c)
		break
	}
}
