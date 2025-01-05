package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)

	totalMinutes := seconds / 60
	h := totalMinutes / 60
	m := totalMinutes % 60

	fmt.Printf("It is %v hours %v minutes.\n", h, m)
}
