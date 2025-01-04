package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 2, 5, 6, 8, 9, 4, 4, 2, 1, 3, 3, 0, 6, 0}
	fmt.Println(retArr(arr))

	r := struct{ n int }{1} //Если тип структуры используется только для одного значения, нам не нужно давать ему имя. Значение может иметь анонимный тип структуры.
	r = struct{ n int }{2}
	r.n = 3
	fmt.Println(r, "r")
	dog := struct {
		name   string
		isGood bool
	}{
		"Aqtos",
		true,
	}
	fmt.Println(dog.name, dog)
	dog.name = "Klyk"
	dog.isGood = false
	fmt.Println(dog.name, dog)
}

func retArr(arr []int) []int {
	res := make([]int, 0, len(arr))
	myMap := make(map[int]struct{}, len(arr)) // Использование struct{} вместо int: Это уменьшит использование памяти в карте, так как пустая структура занимает 0 байт.
	for _, n := range arr {
		if _, ok := myMap[n]; !ok {
			myMap[n] = struct{}{}
			res = append(res, n)
		}
	}
	return res
}
