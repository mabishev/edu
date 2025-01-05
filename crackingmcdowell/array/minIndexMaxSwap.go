package main

import "fmt"

func main() {
	arr := []int{5, 6, 2, 9, 0, 14, 64, 22, 12}
	min := minIndex(arr)
	fmt.Println(arr[min])
	max := maxIndex(arr)
	fmt.Println(arr[max])
	swap(arr, min, max)
	fmt.Println(arr)
}

func minIndex(arr []int) int {
	var minIndex int
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func maxIndex(arr []int) int {
	var maxIndex int
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func swap(arr []int, min, max int) {
	arr[min], arr[max] = arr[max], arr[min]
}
