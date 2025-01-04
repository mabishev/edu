package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := strings.Split(scanner.Text(), " ")
	arr := []int{}

	for _, val := range s {
		n, _ := strconv.Atoi(val)
		arr = append(arr, n)
	}

	fmt.Println(findX(arr, 4))
	fmt.Println(lastX(arr, 2))
}

// находит первое(левое) вхождение числа Х
func findX(arr []int, x int) int {
	for i := range arr {
		if arr[i] == x {
			return i + 1
		}
	}
	return -1
}

// находит последнее(правое) вхождение числа Х
func lastX(arr []int, x int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == x {
			return i + 1
		}
	}
	return -1

	//ans := -1
	//for i := range arr {
	//	if arr[i] == x {
	//		ans = i + 1
	//	}
	//}
	//return ans
}
