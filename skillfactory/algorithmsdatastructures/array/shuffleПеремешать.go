package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(ar), func(i, j int) { ar[i], ar[j] = ar[j], ar[i] })
	fmt.Println(ar)

	//алгоритм без испльзования памяти и локации локальных перменных(поменять местами  0 и 1 элемент)
	ar2 := []int{10, 15, 3, 5, 7}
	ar2[0] = ar2[0] + ar2[1]
	ar2[1] = ar2[0] - ar2[1]
	ar2[0] = ar2[0] - ar2[1]
	fmt.Println(ar2)
}
