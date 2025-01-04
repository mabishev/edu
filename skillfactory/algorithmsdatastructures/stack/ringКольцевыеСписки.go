package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем кольцо из пяти элементов
	r := ring.New(5)
	n := r.Len()
	// Инициализируем элементы кольца значениями
	for i := 0; i < n; i++ {
		//var n int
		//fmt.Scan(&n)
		r.Value = i
		r = r.Next()
	}
	// Перемещаемся вперёд по кольцу и печатаем значение каждого
	// элемента
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	for k := n; k > 0; k-- {
		r = r.Prev()
		fmt.Println(r.Value)
	}
	fmt.Println("///////////////////////////")
	// Создаем кольцо из пяти элементов
	r2 := ring.New(5)
	// Инициализируем элементы кольца значениями
	// Обратите внимание мы обошли кольцо полностью один раз!
	// Полностью!
	for i := 0; i < r2.Len(); i++ {
		r2.Value = i
		r2 = r2.Next()
	}
	for l := 0; l < r2.Len(); l++ {
		fmt.Println(r2.Value)
		r2 = r2.Next()
	}
	// Теперь кто-то из вас может подумать, что повторная
	// итерация по
	// элементам приведет к ошибке
	// ведь мы уже в конце кольца, куда нам дальше-то двигаться?

	// Из конца списка мы вновь возвращаемся автоматически в
	// начало и
	// перемещаемся вперед по кольцу на два элемента
	fmt.Println("///////////////////////////")
	for j := 0; j < 2; j++ {
		r2 = r2.Next()
		fmt.Println(r2.Value)
	}
	// А теперь внимание!
	// Мы создаем новое кольцо - просто указав
	// ссылку на следующий элемент уже существующего кольца
	newRing := r2.Next() // содержит элмементы: 3 4 0 1 2
	r2 = r2.Next()
	fmt.Printf("%d %d\n", newRing.Value, r2.Value)
	for l := 0; l < newRing.Len(); l++ {
		fmt.Println(newRing.Value)
		newRing = newRing.Next()
	}
	fmt.Println("///////////////////////////")

	// Создаем два кольца, каждый размером по 3 элемента
	ring1 := ring.New(3)
	ring2 := ring.New(3)

	// Получаем длины каждого из колец
	ring1Len := ring1.Len()
	ring2Len := ring2.Len()

	// Заполняем первое кольцо нулями
	for i := 0; i < ring1Len; i++ {
		ring1.Value = 0
		ring1 = ring1.Next()
	}

	// Заполняем второе кольцо единичками
	for j := 0; j < ring2Len; j++ {
		ring2.Value = 1
		ring2 = ring2.Next()
	}

	// Соединяем второе кольцо с первым вот таким образом
	ring3 := ring1.Next().Link(ring2)

	// Проверяем содержимое нового кольца
	for i := 0; i < ring3.Len(); i++ {
		fmt.Println(ring3.Value)
		ring3 = ring3.Next()
	}
}
