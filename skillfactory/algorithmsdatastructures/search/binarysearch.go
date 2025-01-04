package main

import "fmt"

func main() {
	fmt.Println(findBinary([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(findBinaryR([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 0, 11))
}

// У алгоритма бинарного поиска логарифмическая сложность — O(log2(N))
// Бинарный поиск используется для нахождения элемента в уже отсортированном слайсе/массиве.
// Его преимущество перед линейным поиском в скорости выполнения.
func findBinary(s []int, val int) int {
	start, end := 0, len(s)
	for {
		middle := (start + end) / 2
		if s[middle] == val {
			return middle
		}
		if middle == start {
			return -1
		}
		if s[middle] < val {
			start = middle
		}
		if s[middle] > val {
			end = middle
		}
	}
}

func findBinaryR(s []int, val, start, end int) int { // рекурсивная версия
	middle := (start + end) / 2
	if s[middle] == val {
		return middle
	}
	if s[middle] == start {
		return -1
	}
	if s[middle] > val {
		return findBinaryR(s, val, start, middle)
	}
	if s[middle] < val {
		return findBinaryR(s, val, middle, end)
	}
	return -1
}
