package edu

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func Rev() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := []int{5, 6, 7, 8, 9, 10}
	reverse(a[:]) //если передать просто массив а, то missmatch types
	reverse(b)
	fmt.Println(a, b)

	c := []int{0, 1, 2, 3, 4, 5}
	// Циклический сдвиг влево на две позиции.
	reverse(c[:2])
	reverse(c[2:])
	reverse(c)
	fmt.Println(c)
}
