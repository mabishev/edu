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
	fmt.Println(findMax2(arr))
}

// находит максимальное число в последовательности и второе по величине число
func findMax2(arr []int) (max, max2 int) {
	max = arr[0]
	max2 = arr[1]
	if max2 > max {
		max, max2 = max2, max
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] > max {
			max2 = max
			max = arr[i]
		}
	}
	return max, max2
}
