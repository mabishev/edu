package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Задача: Дана строка S. Вывести гистограмму как в примере(коды символов отсортированы): S = Hello, world!
//             #
//             ##
//       ##########

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	printChart(s)
}

func printChart(s string) { //рассчитать диаграмму
	symCount := make(map[string]int)
	maxSymCount := 0

	for _, sym := range s {
		symCount[string(sym)]++
		if symCount[string(sym)] > maxSymCount {
			maxSymCount = symCount[string(sym)]
		}
	}

	var arr []string
	for sym := range symCount {
		arr = append(arr, sym)
	}

	sort.Strings(arr)

	for i := maxSymCount; i >= 1; i-- {
		for _, sym := range arr {
			if symCount[sym] >= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
