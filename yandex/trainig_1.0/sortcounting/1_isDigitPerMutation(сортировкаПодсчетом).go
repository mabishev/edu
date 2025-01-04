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
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	fmt.Println(isDigitPerMutation(x, y))
}

// дано два числа без ведущих нолей(0221). Можно ли получить первое из второго перестановкой цифр
// 2021 == 2210 (true);  2021 == 2220 (false)
func isDigitPerMutation(x, y int) bool {
	arrX := countDigits(x)
	arrY := countDigits(y)
	for i := range arrX {
		if arrX[i] != arrY[i] {
			return false
		}
	}
	return true
}

func countDigits(n int) []int {
	arr := make([]int, 10)
	for n > 0 {
		lastDigit := n % 10
		arr[lastDigit] += 1
		n /= 10
	}
	return arr
}
