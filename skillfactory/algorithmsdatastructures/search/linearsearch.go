package main

import "fmt"

func main() {
	fmt.Println(findInSlice([]int{4, 63, 234, 634, 2, 36, 6, 3, 23}, 36))
}

//Алгоритм линейного поиска имеет сложность O(n), так как подразумевает, что надо
//пройтись по всему слайсу/массиву/списку один раз, проверив (n) элементов.

// поиск наибольшего элемента в слайсе целых чисел
func findMaxInSlice(s []int) (int, int) {
	var index, val int
	for i, v := range s {
		if v > val {
			index, val = i, v
		}
	}
	return index, val
}

// Нахождение индекса первого элемента, равного конкретному значению:
func findInSlice(s []int, value int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return i
		}
	}
	return -1
}
