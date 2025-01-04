package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Задача: Дана последовательсность чисел длинной N и M запросов. Запросы: "Сколько нулей на
//полуинтервале [L, R): prefisZeroes[R] - prefixZeroes[L]"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text()) // input: 1 0 1 1 0 0 1
		arr = append(arr, k)
	}
	result := makePrefixZeroes(arr)
	fmt.Println(result, "////")
	fmt.Println(countZeroes(result, 1, 7))
}
func makePrefixZeroes(arr []int) []int { //output: 0 0 1 1 1 2 3 3
	prefixZeroes := make([]int, len(arr)+1)
	for i := 1; i < len(arr)+1; i++ {
		if arr[i-1] == 0 {
			prefixZeroes[i] = prefixZeroes[i-1] + 1
		} else {
			prefixZeroes[i] = prefixZeroes[i-1]
		}
	}
	return prefixZeroes
}

func countZeroes(prefixZeroes []int, l, r int) int {
	return prefixZeroes[r] - prefixZeroes[l]
}
