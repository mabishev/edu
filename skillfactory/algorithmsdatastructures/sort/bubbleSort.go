package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	ar := make([]int, 100000)
	for i := range ar {
		ar[i] = rand.Intn(10000)
	}

	//fmt.Println(ar)
	//bubbleSortSortInts(ar)
	bubbleSortFromGPT(ar)
	//bubbleSort(ar)
	//fmt.Println(ar)
}

// при длинне массива 100000 скорость работы sort.Ints в 500 раз быстрее чем моя функция bubbleSort
func bubbleSortSortInts(ar []int) []int {
	startTime := time.Now()
	sort.Ints(ar) //внутренняя реализация сортировки quick sort - быстрая сортировка

	elapsedTime := time.Since(startTime) // Вычисляем время выполнения
	fmt.Println(elapsedTime)
	return ar
}

func bubbleSortRecursive(ar []int) []int {
	n := len(ar)
	if n <= 1 {
		return ar
	}
	// Проходим по массиву и проверяем пары соседних элементов
	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			// Обмениваем элементы местами, если текущий элемент больше следующего
			ar[i], ar[i+1] = ar[i+1], ar[i]
		}
	}

	// Рекурсивно вызываем функцию для подмассива, исключая последний элемент
	return append(bubbleSortRecursive(ar[:n-1]), ar[n-1])

}
func bubbleSortRecursiveSkillfactory(ar []int) {
	if len(ar) == 1 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i+1], ar[i] = ar[i], ar[i+1]
		}
	}
	bubbleSortRecursive(ar[:len(ar)-1])
}
func bubbleSort(ar []int) []int {
	startTime := time.Now()

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[i] < ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime) // Вычисляем время выполнения
	fmt.Println(elapsedTime)
	return ar
}

// проверка
func bubbleSortFromGPT(ar []int) []int {
	startTime := time.Now()
	n := len(ar)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ { //j < n-i-1, чтобы исключить уже отсортированные элементы на каждой итерации внешнего цикла.
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swapped = true
			}
		}
		if !swapped {
			// Массив уже отсортирован
			break
		}
	}
	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime) // Вычисляем время выполнения
	fmt.Println(elapsedTime)
	return ar
}

func bubbleSortWithInterruptSkillfactory(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		swapped := false
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				swapped = true
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
		if !swapped {
			break
		}
	}
}
