package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		arr = append(arr, i)
	}
	num1, num2 := findTwoNumbersWithSum(arr, n)
	fmt.Println(num1, num2)

}

// находит сумму двух чисел равную х
func findTwoNumbersWithSum(arr []int, x int) (int, int) {
	// Создаем хэш-таблицу для хранения чисел и их индексов. Сложность O(N). Этот вариант чуть эффективнее,
	// т.к. поиск элемента по индексу в слайсе (arr[index]) выполняется за O(1) времени
	numbers := make(map[int]int)
	for i, num := range arr {
		diff := x - num
		if index, found := numbers[diff]; found { // Проверяем, есть ли такая разница в хэш-таблице.
			return arr[index], num
		}
		numbers[num] = i // Если нет, добавляем текущее число в хэш-таблицу.
	}
	return 0, 0

	// Второй вариант Сложность O(N). Поиск элемента по значению в слайсе (num) занимает O(N) времени.
	//numbers := make(map[int]bool)
	//for _, num := range arr {
	//	diff := x - num
	//	if numbers[diff] {
	//		return diff, num
	//	}
	//	numbers[num] = true
	//}
	//return 0, 0
}
