package main

import "fmt"

func main() {
	ar := []int{4, 3, 6, 7, 2, 1, 0}
	ar2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(checkSliceIsSorted(ar))
	fmt.Println(checkSliceIsSorted(ar2))

}
func checkSliceIsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
