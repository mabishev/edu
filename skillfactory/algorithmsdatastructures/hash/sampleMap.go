package main

import "fmt"

//func main() {
//	var f, s int
//	myMap := make(map[int]int)
//	fmt.Println("Enter first array size:")
//	fmt.Scan(&f)
//	fmt.Println("Enter second array size:")
//	fmt.Scan(&s)
//	var ar1, ar2 int
//	fmt.Println("Enter first array:")
//	for i := 0; i < f; i++ {
//		fmt.Scan(&ar1)
//		myMap[i] = ar1
//	}
//	var sl []int
//	fmt.Println("Enter second array:")
//	for j := 0; j < s; j++ {
//		fmt.Scan(&ar2)
//		for k := 0; k < f; k++ {
//			if myMap[k] == ar2 {
//				sl = append(sl, ar2)
//			}
//		}
//	}
//	fmt.Println(sl)
//
//}

type user struct {
	name, login, email string
	age                int
}
type userString string

func main() {
	var hashMap = map[string]user{
		"Dolly": {name: "Mrs Dolly", login: "HelloDolly", email: "holly@gmail.com", age: 33},
	}

	// Инициализируем переменную с использованием функции make
	// В качестве первого аргумента передаем тип, а авторым указываем ожидаемый размер.
	// Благодаря второму аргументу, go сгенерирует хеш-функцию, оптимизированную на 2000 элемнтов
	makedHashMap := make(map[userString]user, 2000)
	fmt.Println(hashMap)
	fmt.Println(makedHashMap)
	fmt.Println(len(hashMap))
	// Мы можем перебрать все элементы мапы в цикле
	// ВНИМАНИЕ. Go __не гарантирует__ порядок. Т.е. очередность получения элементов будет случайной
	for identity, u := range hashMap {
		fmt.Printf("[%s] %s's email is %s", identity, u.name, u.email)
	}
}
