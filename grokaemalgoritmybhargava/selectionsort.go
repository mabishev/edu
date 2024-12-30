package edu

import "fmt"

func test() {
	arr := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(arr))
}

// сложность O(N2) (N в квадрате)
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := []int{}
	for i := 0; i < len(arr); i++ {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
		i--
	}
	return newArr
}
