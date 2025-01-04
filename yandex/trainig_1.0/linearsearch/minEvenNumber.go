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
	fmt.Println(minEvenNumber(arr))
}

// находит минимальное четное число
func minEvenNumber(arr []int) int {
	chet := -1
	for _, val := range arr {
		if val%2 == 0 && (chet == -1 || val < chet) { // val > chet - макс четное число
			chet = val
		}
	}
	return chet
}
