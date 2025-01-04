package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Задача: Префиксные суммы. Массив nums из N чисел, чему равна сумма элементов на полуинтервале[L, R).
// Подсчитаем массив prefixSum длиной N+1, где prefixSum[K] будет хранить сумму всех чисел из nums с
//индексами от 0 до К-1. Массив можно построить за O(N): prefixSum[i] = prefix[i-1]+nums[i-1]. Массив
//prefixSum должен быть длиннее nums на 1. Ответ на запрос суммы на отрезке за О(1): sum(L,R) =
//prefiSum[R]-prefixSum[L]. Найдем sum(2, 5) = prefixSum[5]-prefixSum[2] = 21 - 8 = 13

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text()) // input: 5 3 8 1 4 6
		arr = append(arr, k)
	}
	result := makePrefixSum(arr)

	fmt.Println(RSQ(result, 2, 5))
}
func makePrefixSum(arr []int) []int { //output: 0 5 8 16 17 21 27
	prefixSum := make([]int, len(arr)+1)
	for i := 1; i < len(arr)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
	}
	return prefixSum
}

// реализация RSQ(range sum query) через префиксные суммы
func RSQ(prefixSum []int, l, r int) int { // output: 13
	return prefixSum[r] - prefixSum[l]
}
