package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	RUR
)

func main() {
	q := [...]string{RUR: "P", USD: "$", EUR: "Eu"} //В этом случае индексы могут появляться в любом порядке, а некоторые из них мо­ гут быть опущены
	s := [...]int{10: 99}                           //Если в литерале массива на месте длины находится троеточие “ . . . ”, то длина массива определяется количеством инициализаторов.
	fmt.Println(q, s)

	fmt.Println(out(nil, 5))            // можно вместо nil вставить []int{}
	fmt.Printf("%q\n", []rune("Салем")) //встроенное преобразование
}

func out(stack []int, n int) []int {
	for i := 0; i < n; i++ {
		stack = append(stack, i)
	}
	return stack
}
