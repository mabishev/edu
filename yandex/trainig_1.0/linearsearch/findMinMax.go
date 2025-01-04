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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	start := time.Now()

	arr := []int{}
	s := strings.Split(scanner.Text(), " ")

	for _, val := range s {
		n, _ := strconv.Atoi(val)
		arr = append(arr, n)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
	fmt.Println(findMin(arr))
	fmt.Println(findMax(arr))
}

func findMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
