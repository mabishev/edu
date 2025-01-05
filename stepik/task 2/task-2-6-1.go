package main

import (
	"fmt"
)

func divide(x, y int) (int, error) {
	return x / y, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	reSult, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(reSult)
	}
}
