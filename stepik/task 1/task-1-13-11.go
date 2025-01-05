package main

import "fmt"

func main() {
	var a, b, max int

	for fmt.Scan(&a, &b); b >= a; b-- {
		if b%7 == 0 {
			max = b
			fmt.Print(max)
			break
		} else if b == a {
			fmt.Println("NO")
		}

	}
}
