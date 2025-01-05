package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TaskB() {
	//var t int   // медленнее чем bufio.NewScanner()  5 баллов
	//fmt.Scan(&t)
	//for i := 0; i < t; i++ {
	//	var n int
	//	fmt.Scan(&n)
	//	myMap := make(map[int]int)
	//	for j := 0; j < n; j++ {
	//		var p int
	//		fmt.Scan(&p)
	//		myMap[p]++
	//	}

	scanner := bufio.NewScanner(os.Stdin) // этот код работает быстрее чем fmt.Scan() 10 баллов
	scanner.Split(bufio.ScanWords)        // одной строкой через пробел ввести цифры (напр. 2 4 6)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		myMap := make(map[int]int)
		for j := 0; j < n; j++ {
			scanner.Scan()
			p, _ := strconv.Atoi(scanner.Text())
			myMap[p]++
		}

		var sum int
		for key, val := range myMap {
			sum += key * (val - val/3)
		}
		fmt.Println(sum)

	}

}
