package main

import (
	"fmt"
)

func main() {
	setSize := 10
	arr := make([][]int, setSize)       // хэш-таблица
	add(arr, 86, 96, 76, 6, 16, 26, 36) // совпадение значений хэш-функции для разных параметров - коллизия
	add(arr, 27)
	fmt.Println(arr)
	fmt.Println(find(arr, 16))
	fmt.Println(delete(arr, 16), "delete")
}

func add(arr [][]int, x ...int) [][]int {
	for _, val := range x {
		arr[val%10] = append(arr[val%10], val) // хеш-функция это - F(X) = X % setSize
	}
	return arr
}

// удаляет необходимое значение путем замены последнним элементом и удалением последнего элемента
func delete(arr [][]int, x int) [][]int {
	needIndex := x % 10
	for i := 0; i < len(arr[needIndex]); i++ {
		if arr[needIndex][i] == x {
			arr[needIndex][i] = arr[needIndex][len(arr[needIndex])-1]
			arr[needIndex] = arr[needIndex][:len(arr[needIndex])-1]
			break
		}
	}
	return arr
}

func find(arr [][]int, x int) bool {
	for _, val := range arr[x%10] {
		if val == x {
			return true
		}
	}
	// бинарный поиск
	//needIndex := x % 10
	//sort.Ints(arr[needIndex])
	//first := arr[needIndex][0]
	//last := arr[needIndex][len(arr[needIndex])-1]
	//for first <= last {
	//	mid := (first + last) / 2
	//	guess := arr[needIndex][mid]
	//	if guess == x {
	//		return true
	//	}
	//	if guess > x {
	//		last = mid - 1
	//	} else {
	//		first = mid + 1
	//	}
	//
	//}
	return false
}
