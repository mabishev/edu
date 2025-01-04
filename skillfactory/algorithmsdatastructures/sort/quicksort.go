package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	quickSort(ar)

	fmt.Println(ar)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr)) // Выбираем случайный опорный элемент

	arr[pivot], arr[right] = arr[right], arr[pivot] // Перемещаем опорный элемент в конец массива

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i] // Перемещаем меньшие элементы влево
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left] // Возвращаем опорный элемент на свое место

	quickSort(arr[:left])   // Рекурсивно сортируем левую часть
	quickSort(arr[left+1:]) // Рекурсивно сортируем правую часть
}

func quickSortSkillfactory(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])

	return
}
