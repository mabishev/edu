package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 4)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	n := len(ar)
	for i := 1; i < n; i++ {
		key := ar[i]
		j := i - 1

		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j--
		}

		ar[j+1] = key
	}
}

func InsertionSortSkillfactory(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}
