package edu

import (
	"fmt"
)

//Этот алгоритм нахо­дит путь с минимальным количеством сегментов. Невзвешенный граф

// Функция для проверки, является ли человек продавцом манго
func personIsSeller(person string) bool {
	// Здесь может быть ваша логика для определения продавца манго
	return person[len(person)-1] == 'm'
}

// Функция поиска продавца манго
func search(name string) bool {
	// Граф друзей
	var graph = map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"alice":  {"peggy"},
		"bob":    {"anuj", "peggy"},
		"claire": {"thom", "alice"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	searchQueue := []string{}
	searchQueue = append(searchQueue, graph[name]...)
	//searched := []string{}
	searched := make(map[string]int)

	// с использованием карты
	for len(searchQueue) != 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]

		if _, ok := searched[person]; !ok {
			if personIsSeller(person) {
				fmt.Println(person + " is a mango seller")
				return true
			} else {
				searchQueue = append(searchQueue, graph[person]...)
				searched[person]++
			}
		}

		// С использованием массива
		// Проверяем, не искали ли мы уже этого человека
		// alreadySearched := false
		// for _, searchedPerson := range searched {
		// 	if person == searchedPerson {
		// 		alreadySearched = true
		// 		break
		// 	}
		// }

		// if !alreadySearched {
		// 	if personIsSeller(person) {
		// 		fmt.Println(person + " is a mango seller")
		// 		return true
		// 	} else {
		// 		searchQueue = append(searchQueue, graph[person]...)
		// 		searched = append(searched, person)
		// 	}
		// }
	}
	return false
}

func main() {
	found := search("you")
	if !found {
		fmt.Println("No mango seller found")
	}
	// startPesron := "you" // для использования searchList
	// if !searchList(graph, startPesron) {
	// 	fmt.Println("No mango seller found")
	// }
}

// пакет container/list предоставляет двунаправленный связанный список. Очередь list используется для хранения и обработки элементов в порядке их добавления.
// func searchList(graph map[string][]string, start string) bool {
// 	searchQueue := list.New()
// 	for _, friend := range graph[start] { // Добавляем друзей стартовой персоны в очередь
// 		searchQueue.PushBack(friend)
// 	}
// 	searched := make(map[string]bool)

// 	for searchQueue.Len() > 0 {
// 		person := searchQueue.Remove(searchQueue.Front()).(string) // Извлекаем первого человека из очереди. (string) приводим тип any к типу string
// 		if _, ok := searched[person]; !ok {
// 			if personIsSeller(person) {
// 				fmt.Println(person, "является продавцом манго!")
// 				return true
// 			} else {
// 				for _, friends := range graph[person] {
// 					searchQueue.PushBack(friends)
// 				}
// 				searched[person] = true // Отмечаем, что текущий человек уже был проверен
// 			}
// 		}
// 	}
// 	return false // Продавец манго не найден
// }
