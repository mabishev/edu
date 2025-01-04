package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	ar := make([]int, 10)
	ar2 := []int{1, -1}
	for i := range ar {
		ar[i] = rand.Intn(100)
	}
	if !checkSliceIsSorted2(ar2) {
		quisckSort(ar2)
		fmt.Println(ar2)
	}

}

func quisckSort(ar []int) []int {
	sort.Ints(ar)
	return ar
}
func checkSliceIsSorted2(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
