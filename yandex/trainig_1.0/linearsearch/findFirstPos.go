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
	fmt.Println(first(arr, 2))
}

func first(arr []int, x int) int {
	posit := make(map[int]int)
	for i, val := range arr {
		if pos, found := posit[val]; found && val == x {
			return pos + 1
		}
		posit[val] = i

	}
	return -1
}
