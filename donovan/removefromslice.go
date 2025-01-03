package edu

import "fmt"

// удалить элемент из средины среза, сохранив порядок оставшихся элементов
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Если нам не обязательно сохранять порядок, можно просто перенести последний элемент на место удаляемого:
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Rem() {
	stack := []int{}

	//поместить новое значение в стек
	for i := 1; i < 10; i++ {
		stack = append(stack, i) // внесение i в stack
	}
	//Вершина стека представляет собой последний его элемент:
	top := stack[len(stack)-1]

	//Снятие элемента со стека:
	stack = stack[:len(stack)-1] //  Удаление элемента из стека

	fmt.Println(top)
	stack = remove(stack, 3)
	fmt.Println(stack, "remove")

	stack = remove2(stack, 3)
	fmt.Println(stack, "remove2")
}
