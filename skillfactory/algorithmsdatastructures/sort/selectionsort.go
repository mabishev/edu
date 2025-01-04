package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	doubleSort(ar)

	fmt.Println(ar)

}

func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minInd := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minInd] {
				minInd = j
			}
		}
		ar[i], ar[minInd] = ar[minInd], ar[i] // Обмен значениями
	}
}

// работает «справа налево» (поиск максимальных элементов и перемещение их в конец).
func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		maxInd := i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxInd] {
				maxInd = j
			}
		}
		ar[maxInd], ar[i] = ar[i], ar[maxInd]
	}
}

// двунаправоенная сортировка
func doubleSort(ar []int) { //13, 6, 11, -5, 9, 0
	for i := 0; i < len(ar); i++ {
		min := i
		max := len(ar) - i - 1
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[max] < ar[j] {
				max = j
			}
			if ar[min] > ar[j] {
				min = j
			}
		}
		ar[max], ar[min] = ar[min], ar[max]
		ar[min], ar[len(ar)-i-1] = ar[len(ar)-i-1], ar[min]
	}
}
