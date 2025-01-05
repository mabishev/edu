package main

import "fmt"

func work(x int) int {
	if x < 3 {
		return x - 1
	}
	return x + 1
}

func main() {
	var a int
	var m = map[int]int{}
	for i := 1; i < 11; i++ {
		fmt.Scan(&a)
		if _, ok := m[a]; !ok {
			m[a] = work(a)
		}
		fmt.Print(m[a], " ")
	}

}
