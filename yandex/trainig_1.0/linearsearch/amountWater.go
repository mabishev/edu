package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // aa b cc d
	scanner.Scan()

	start := time.Now()

	s := strings.Split(scanner.Text(), " ")
	arr := []int{}
	for _, val := range s {
		n, _ := strconv.Atoi(val)
		arr = append(arr, n)
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println(elapsed)
	fmt.Println(water(arr))
}

// находит количество воды
func water(arr []int) int { // пример входных данных 3 1 4 3 5 1 1 3 1 (ответ 7)
	maxPos := 0

	for i := range arr {
		if arr[i] > arr[maxPos] {
			maxPos = i
		}
	}

	var total, nowMax int

	for i := 0; i <= maxPos; i++ {
		if arr[i] > nowMax {
			nowMax = arr[i]
		}
		total += nowMax - arr[i]
	}

	nowMax = 0

	for i := len(arr) - 1; i > maxPos; i-- {
		if arr[i] > nowMax {
			nowMax = arr[i]
		}
		total += nowMax - arr[i]
	}
	return total
}
